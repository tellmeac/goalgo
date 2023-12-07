package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tellmeac/goalgo/internal/app"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	ctx := context.Background()

	token := os.Getenv("BOT_TOKEN")

	channelChatID, err := strconv.ParseInt(os.Getenv("PUBLISH_CHANNEL"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Fatal(run(ctx, bot, channelChatID))
}

func run(ctx context.Context, bot *tgbotapi.BotAPI, chatID int64) error {
	offset := time.Now().Unix()

	for {
		chart, err := poll(ctx, offset)
		if err != nil {
			log.Printf("poll err: %s", err.Error())
			continue
		}

		offset = chart.Data[len(chart.Data)-1].Time

		msg := &tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID: chatID,
			},
			Text: fmt.Sprintf("chart update: %d", len(chart.Data)), // TODO: render data
		}

		if _, err := bot.Send(msg); err != nil {
			log.Printf("send err: %s", err.Error())
		}
	}
}

func poll(ctx context.Context, from int64) (app.Chart, error) {
	for {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://localhost:8080/updates?from=%d", from), nil)
		if err != nil {
			cancel()
			return app.Chart{}, err
		}

		resp, err := http.DefaultClient.Do(req)
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("keep polling...")
			continue
		}
		if err != nil {
			cancel()
			return app.Chart{}, err
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			cancel()
			_ = resp.Body.Close()
			return app.Chart{}, err
		}

		var result app.Chart
		if err := json.Unmarshal(data, &result); err != nil {
			cancel()
			_ = resp.Body.Close()
			return app.Chart{}, err
		}

		cancel()
		_ = resp.Body.Close()

		return result, nil
	}
}
