package v1

import (
	"github.com/gin-gonic/gin"

	posthttp "routing/internal/transport/http/v1/posts"
)

type Handlers struct {
	Posts *posthttp.Handler
}

func RegisterRoutes(router *gin.RouterGroup, h Handlers) {
	h.Posts.RegisterRoutes(router.Group("/posts"))
}
