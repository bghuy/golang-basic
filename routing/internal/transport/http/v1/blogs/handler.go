package blogs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("", h.List)
	rg.POST("", h.Create)
	rg.GET("/:id", h.GetByID)
}

func (h *Handler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of blogs"})
}

func (h *Handler) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"title": c.PostForm("title")})
}

func (h *Handler) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
}
