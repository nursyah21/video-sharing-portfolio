package routes

import (
	"video-go/database"
	"video-go/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		h := handlers.NewHandler(database.Db)
		api.GET("/videos", h.GetVideos)
		api.POST("/videos", h.PostVideo)
		api.PUT("/videos/:id", h.UpdateVideo)
		api.DELETE("/videos/:id", h.DeleteVideo)
	}
}
