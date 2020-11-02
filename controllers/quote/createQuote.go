package quote

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"quote-api/core/models"
	"time"
)

type CreateRequestBody struct {
	Author string `form:"author" json:"author" binding:"required"`
	Quote  string `form:"quote" json:"quote" binding:"required"`
}

type CreateResponseBody struct {
	QuoteId interface{} `json:"quoteId"`
}

/*** Create Quote Handle
 * POST - /quote
 * REQ:{BODY: author,quote :string :required} RES:{BODY: quoteId :string}
 */

func CreateQuote(c *gin.Context) {
	body := CreateRequestBody{}
	if c.ShouldBindBodyWith(&body, binding.JSON) != nil {
		c.Status(422)
		return
	}
	quote := models.Quote{
		Author:    body.Author,
		Quote:     body.Quote,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := QuotesCollection.InsertOne(context.TODO(), quote)
	log.Println(res)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotModified)
	} else {
		log.Println(res.InsertedID)
		c.JSON(http.StatusOK, &CreateResponseBody{
			res.InsertedID,
		})
	}
}
