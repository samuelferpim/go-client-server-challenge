package main

import (
	"log"
	"net/http"

	"github.com/samuelferpim/go-client-server-challenge/server/internal/databases"
	"github.com/samuelferpim/go-client-server-challenge/server/internal/routes"
)

func main() {
	log.Println("starting server...")
	log.Println("preparing database...")
	db, err := databases.Setup()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("migrating database...")
	err = databases.Migrate(db)
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/price", routes.ExchangeRouteHandler)
	http.HandleFunc("/exchange/history", routes.ExchangeHistoryHandler)
	http.ListenAndServe(":8080", nil)
}
