package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tellmeac/goalgo/internal/app"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func buildMessage(chatID int64, chart app.Chart) (*tgbotapi.MessageConfig, bool) {
	var last *app.Stamp
	for idx := range chart.Data {
		s := &chart.Data[idx]

		if s.NeedPoint == nil {
			continue
		}

		last = s
	}

	if last == nil {
		return nil, false
	}

	args := struct {
		ClosePrice float64
		StopLoss   float64
		TakeProfit float64
	}{
		ClosePrice: last.Candlestick.Close,
		StopLoss:   last.Candlestick.Close - (last.TopLine - last.DownLine),
		TakeProfit: last.Candlestick.Close + (last.TopLine - last.DownLine),
	}

	t, err := template.New("message-body").Parse(`Обнаружена благоприятная точка для покупки: 
Цена {{ printf "%.3f" .ClosePrice }}
Рекомендованный StopLoss: {{ printf "%.3f" .StopLoss }}
Рекомендованный TakeProfit: {{ printf "%.3f" .TakeProfit }}`)
	if err != nil {
		log.Fatal(err)
	}

	result := &strings.Builder{}
	if err := t.Execute(result, args); err != nil {
		log.Fatal(err)
	}

	return &tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: chatID,
		},
		Text: result.String(),
	}, true
}

func run(ctx context.Context, bot *tgbotapi.BotAPI, publishChatID int64) error {
	offset := time.Now().Unix()
	// TODO: debug purposes
	// offset := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC).Unix()

	for {
		chart, err := poll(ctx, offset)
		if err != nil {
			log.Printf("poll err: %s", err.Error())
			continue
		}

		offset = chart.Data[len(chart.Data)-1].Time

		msg, ok := buildMessage(publishChatID, chart)
		if !ok {
			continue
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
