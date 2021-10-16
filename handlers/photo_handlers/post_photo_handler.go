package photo_handlers

import (
	"github.com/m0thm4n/UploadService/domain/photo"
	"github.com/m0thm4n/UploadService/handlers/request_models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var PostSinglePhotoHandler = func(r *photo.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rm *request_models.CreatePhotoModel

		err := c.ShouldBindJSON(&rm)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not bind the request!"})
			return
		}

		photo := &photo.Photo{
			DateTimeOriginal: rm.DateTimeOriginal,
			Manufacturer: rm.Manufacturer,
			Model: rm.Model,
		}

		err = r.Insert(photo)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusAccepted, "Photo Created!")
	}
}