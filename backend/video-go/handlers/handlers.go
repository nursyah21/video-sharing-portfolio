package handlers

import (
	"net/http"
	"video-go/models"
	"video-go/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) PostVideo(c *gin.Context) {
	var video models.Video

	if !utils.BindAndValidate(c, &video) {
		return
	}

	// validate data
	{
		if video.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
			return
		}
		if video.Description == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "description is required"})
			return
		}
		if video.URL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
			return
		}
	}

	h.DB.Create(&video)
	c.JSON(http.StatusCreated, video)
}

func (h *Handler) GetVideos(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		var video models.Video

		if !utils.CheckExist(c, h.DB, "Video not found", &video, id) {
			return
		}

		c.JSON(http.StatusOK, video)
		return
	}

	var videos []models.Video

	// just for benchmark raw speed vs orm
	raw := c.Query("raw")
	if raw == "1" {
		h.DB.Raw("SELECT * FROM videos").Scan(&videos)
	} else {
		h.DB.Find(&videos)
	}

	c.JSON(http.StatusOK, videos)
}

func (h *Handler) UpdateVideo(c *gin.Context) {
	id := c.Param("id")
	var video models.Video

	if !utils.CheckExist(c, h.DB, "Video not found", &video, id) {
		return
	}

	if !utils.BindAndValidate(c, &video) {
		return
	}

	h.DB.Save(&video)
	c.JSON(http.StatusOK, video)
}

func (h *Handler) DeleteVideo(c *gin.Context) {
	id := c.Param("id")

	if !utils.DeleteIfExist(c, h.DB, "Video not found", &models.Video{}, id) {
		return
	}

	c.Status(http.StatusNoContent)
}
