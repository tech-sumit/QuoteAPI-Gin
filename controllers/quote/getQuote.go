package quote

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"quote-api/core/models"
)

/*** Get Quote Handle
 * GET - /quote
 * REQ:{QUERY: search :string} RES:{[author,quote :string]}
 */
func GetQuote(c *gin.Context) {
	query := c.DefaultQuery("search", "")
	var err error
	var res *mongo.Cursor
	if query == "" {
		res, err = QuotesCollection.Find(context.TODO(), &bson.M{})
	} else {
		res, err = QuotesCollection.Find(context.TODO(),
			bson.M{"$text": bson.M{"$search": query}},
			&options.FindOptions{Sort: bson.M{"score": bson.M{"$meta": "textScore"}}})
	}
	log.Println(res)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotFound)
	} else {
		var results []models.Quote
		err := res.All(context.TODO(), &results)
		log.Println(results)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
		} else {
			if len(results) > 0 {
				c.JSON(http.StatusOK, results)
			} else {
				c.Status(http.StatusNoContent)
			}
		}
	}
}
