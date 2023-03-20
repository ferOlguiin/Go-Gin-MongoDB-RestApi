package routes

import (
	"context"
	"golangGinMongo/collection"
	"golangGinMongo/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeletePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DATABASE = database.ConnectDB()
	postId := c.Param("postId")

	var collection = collection.GetCollection(DATABASE, "Posts")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully", "data": res})
}
