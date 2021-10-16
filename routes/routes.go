package routes

import (
	"UploadService/domain/photo"
	"UploadService/handlers/photo_handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(e *gin.Engine, r *photo.Repository) {
	e.GET("/photo/:id", photo_handlers.GetPhotoByIdHandler(r))
	// e.POST("/photo/upload", PostSinglePhotoHandler(r))
	e.DELETE("/photo/delete/:id", photo_handlers.DeleteUploadHandler(r))
	e.POST("/photo/upload", photo_handlers.UploadPhotoHandler(r))
}