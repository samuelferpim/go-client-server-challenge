package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/samuelferpim/go-client-server-challenge/client/internal/model"
)

func GetExchangeRate(ctx context.Context, url string) (*model.ExchangeRate, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var rate model.ExchangeRate
	err = json.Unmarshal(body, &rate)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		err = fmt.Errorf("%d: %s", res.StatusCode, rate.Error)
		return nil, err
	}
	return &rate, nil
}
