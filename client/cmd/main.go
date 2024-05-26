package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/samuelferpim/go-client-server-challenge/client/internal/services"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	url := "http://localhost:8080/price"
	exchange_rate, err := services.GetExchangeRate(ctx, url)
	if err != nil {
		log.Fatalln(err)
	}
	data, err := json.Marshal(exchange_rate)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(data))
	err = services.StoreFile("price.txt", exchange_rate)
	if err != nil {
		log.Fatalln(err)
	}
}
