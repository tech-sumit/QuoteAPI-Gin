package main

import (
	_ "github.com/joho/godotenv/autoload"
	"quote-api/controllers/quote"
	"quote-api/router"
)

func main() {
	//Database Init
	quote.Init()

	//Server Start
	router.Run()
}
