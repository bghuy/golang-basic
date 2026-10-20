package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		posts := v1.Group("/posts")
		{
			posts.GET("/", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"message": "List of posts"})
			})
		}
	}

	router.Run(":8080")
}
