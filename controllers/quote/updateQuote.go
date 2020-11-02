package quote

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"quote-api/core/models"
	"time"
)

type UpdateRequestBody struct {
	ID     string `form:"id" json:"id,omitempty" binding:"required"`
	Author string `form:"author" json:"author,omitempty"`
	Quote  string `form:"quote" json:"quote,omitempty"`
}

/*** Update Quote Handle
 * PUT - /quote
 * REQ:{BODY: id,author,quote :string}
 */

func UpdateQuote(c *gin.Context) {
	body := UpdateRequestBody{}
	if c.ShouldBindBodyWith(&body, binding.JSON) != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	quote := models.Quote{
		Author:    body.Author,
		Quote:     body.Quote,
		UpdatedAt: time.Now(),
	}
	oid, err := primitive.ObjectIDFromHex(body.ID)
	if err != nil {
		log.Println("Can't parse string to ID", err)
		c.Status(http.StatusInternalServerError)
	}
	res, err := QuotesCollection.UpdateOne(context.TODO(),
		bson.M{"_id": oid}, bson.M{"$set": quote})
	log.Println(res)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotModified)
	} else {
		c.Status(http.StatusOK)
	}
}
