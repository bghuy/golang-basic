package v1

import (
	"github.com/gin-gonic/gin"

	bloghttp "routing/internal/transport/http/v1/blogs"
	posthttp "routing/internal/transport/http/v1/posts"
)

type Handlers struct {
	Posts *posthttp.Handler
	Blogs *bloghttp.Handler
}

func RegisterRoutes(router *gin.RouterGroup, h Handlers) {
	h.Posts.RegisterRoutes(router.Group("/posts"))
	h.Blogs.RegisterRoutes(router.Group("/blogs"))
}
