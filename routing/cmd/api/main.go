package main

import (
	"github.com/gin-gonic/gin"

	app "routing/internal/app"
)

func main() {
	router := gin.Default()

	app.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
