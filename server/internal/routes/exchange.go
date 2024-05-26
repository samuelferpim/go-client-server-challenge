package routes

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/samuelferpim/go-client-server-challenge/server/internal/databases"
	"github.com/samuelferpim/go-client-server-challenge/server/internal/model"
	"github.com/samuelferpim/go-client-server-challenge/server/internal/services"
)

func ExchangeRouteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	log.Println("Request started")
	defer log.Println("Request ended")

	rate, err := services.GetExchangeRate(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "context deadline exceeded") {
			err = errors.New("request timed out by server")
		}
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(model.ErrorResponseJson(err))
		return
	}

	data, err := json.Marshal(rate)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(model.ErrorResponseJson(err))
		return
	}

	err = databases.InsertRate(rate)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(model.ErrorResponseJson(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func ExchangeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("request started")
	defer log.Println("request ended")

	rates, err := databases.RateHistory()
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(model.ErrorResponseJson(err))
		return
	}

	data, err := json.Marshal(rates)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(model.ErrorResponseJson(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
