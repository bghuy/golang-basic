package blogs

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

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
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	uploadDir := filepath.Join(root, "uploads")
	dst := filepath.Join(uploadDir, file.Filename)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Upload thành công",
		"filename": file.Filename,
		"size":     file.Size,
	})
}

func (h *Handler) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
}
