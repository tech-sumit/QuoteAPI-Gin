package quote

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type DeleteRequestBody struct {
	ID string `form:"id" json:"id" binding:"required"`
}

/*** Delete Quote Handle
 * DELETE - /quote
 * REQ:{BODY: id :string :required}
 */

func DeleteQuote(c *gin.Context) {
	body := UpdateRequestBody{}
	if c.ShouldBindBodyWith(&body, binding.JSON) != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	oid, err := primitive.ObjectIDFromHex(body.ID)
	if err != nil {
		log.Println("Can't parse string to ID", err)
		c.Status(http.StatusNotModified)
	}
	res, err := QuotesCollection.DeleteOne(context.TODO(), &bson.M{"_id": oid})
	log.Println(res)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotModified)
	} else {
		c.Status(http.StatusOK)
	}
}
