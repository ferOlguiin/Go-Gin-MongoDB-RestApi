package routes

import (
	"context"
	"net/http"
	"time"

	"golangGinMongo/collection"
	"golangGinMongo/database"
	"golangGinMongo/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadOnePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var DATABASE = database.ConnectDB()
	var collectionPost = collection.GetCollection(DATABASE, "Posts")

	postId := c.Param("postId") // "http://localhost:4000/posts/6adh28913271391hduiqg71723" GET METHOD

	var result model.Posts

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)
	err := collectionPost.FindOne(ctx, bson.M{"_id": objId}).Decode(&result)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": result})

}
