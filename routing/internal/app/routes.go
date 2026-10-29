package app

import (
	httpv1 "routing/internal/transport/http/v1"
	posthttp "routing/internal/transport/http/v1/posts"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	httpv1.RegisterRoutes(api.Group("/v1"), httpv1.Handlers{
		Posts: posthttp.NewHandler(), // sau này: NewHandler(postService)
	})
}
