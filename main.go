package main

import (
	"golangGinMongo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/", routes.CreatePost)
	router.GET("/post/:postId", routes.ReadOnePost)
	router.DELETE("/post/:postId", routes.DeletePost)
	router.PUT("/post/:postId", routes.UpdatePost)

	router.Run("localhost:3000")
}
