package routes

import (
	"context"
	"golangGinMongo/collection"
	"golangGinMongo/database"
	"golangGinMongo/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var DATABASE = database.ConnectDB()
	var postCollection = collection.GetCollection(DATABASE, "Posts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	post := new(model.Posts)

	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	postPayload := model.Posts{
		Title:   post.Title,
		Article: post.Article,
	}

	result, err := postCollection.InsertOne(ctx, postPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Posted successfully",
		"Data":    result,
		"item":    postPayload,
	})

}
