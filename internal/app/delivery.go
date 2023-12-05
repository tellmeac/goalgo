package app

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func NewDelivery(s *Service) *Delivery {
	return &Delivery{service: s}
}

type Delivery struct {
	service *Service
}

func (h *Delivery) GetUpdates(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	from := req.URL.Query().Get("from")

	offset, err := strconv.ParseInt(from, 10, 64)
	if err != nil {
		HandleBadRequest(w, err)
		return
	}

	backoff := time.Second
	for {
		chart, err := h.service.GetLatest(ctx, offset)
		if err != nil {
			HandleInternalError(w, err)
			return
		}

		if len(chart.Data) == 0 {
			time.Sleep(backoff)
			backoff = min(time.Minute, 2*backoff)

			continue
		}

		ResponseJSON(w, chart)
		return
	}
}

func (h *Delivery) GetChart(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	fromStr := req.URL.Query().Get("from")
	toStr := req.URL.Query().Get("to")

	from, err := strconv.ParseInt(fromStr, 10, 64)
	if err != nil {
		HandleBadRequest(w, err)
		return
	}

	to, err := strconv.ParseInt(toStr, 10, 64)
	if err != nil {
		HandleBadRequest(w, err)
		return
	}

	chart, err := h.service.GetInPeriod(ctx, from, to)
	if err != nil {
		HandleInternalError(w, err)
		return
	}

	ResponseJSON(w, chart)
}

func ResponseJSON(w http.ResponseWriter, val any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(val)

	_, _ = w.Write(bytes)
}

func HandleBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(err.Error()))
}

func HandleInternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}
