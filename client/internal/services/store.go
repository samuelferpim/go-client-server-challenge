package services

import (
	"fmt"
	"os"

	"github.com/samuelferpim/go-client-server-challenge/client/internal/model"
)

func StoreFile(filename string, exchange *model.ExchangeRate) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintf(file, "Dólar: %s", exchange.Bid)
	return nil
}
