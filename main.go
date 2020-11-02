package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"quote-api/controllers/quote"
	"quote-api/router"
	connections "quote-api/services/mongodb"
)

func main() {
	//Mongo connection Init
	go func() {
		client, err := connections.MongoInit()
		if err != nil {
			log.Fatal(err)
		} else {
			quote.Init(client)
		}
	}()

	//Server Start
	router.Run()
}
