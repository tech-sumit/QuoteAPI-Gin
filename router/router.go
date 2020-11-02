package router

import (
	//"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"strings"

	//"golang.org/x/crypto/acme/autocert"
	"log"
	"os"
	"quote-api/controllers/quote"
)

func Run() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Basic Auth
	authRouter := router.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("AUTH_USERNAME"): os.Getenv("AUTH_PASSWORD"),
	}))

	//Quote Routes
	authRouter.GET("/quote", quote.GetQuote)
	authRouter.PUT("/quote", quote.UpdateQuote)
	authRouter.POST("/quote", quote.CreateQuote)
	authRouter.DELETE("/quote", quote.DeleteQuote)

	//START SERVER
	var host strings.Builder
	host.WriteString(os.Getenv("SERVER_HOST"))
	host.WriteString(":")
	host.WriteString(os.Getenv("PORT"))
	err := router.Run(host.String())
	if err != nil {
		log.Fatal(err)
	}
}
