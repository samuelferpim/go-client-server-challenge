package databases

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/samuelferpim/go-client-server-challenge/server/internal/model"
)

/* Inicia a instância do DB */
func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "exchange_rate.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *sql.DB) error {
	statement, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS exchange_rate (id VARCHAR(36) PRIMARY KEY, rate VARCHAR(16) NULL)",
	)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}

func InsertRate(exchange_rate *model.ExchangeRate) error {
	db, err := Setup()
	if err != nil {
		return err
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, "insert into exchange_rate(id, rate) values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(exchange_rate.Id, exchange_rate.Bid)
	return err
}

func RateHistory() ([]model.ExchangeRate, error) {
	db, err := Setup()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select id, rate from exchange_rate")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rates []model.ExchangeRate
	for rows.Next() {
		var p model.ExchangeRate
		err = rows.Scan(&p.Id, &p.Bid)
		if err != nil {
			return nil, err
		}
		rates = append(rates, p)
	}
	return rates, nil
}
