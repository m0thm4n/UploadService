package photo_handlers

import (
	"UploadService/domain/photo"
	"UploadService/handlers/representation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var GetPhotoByIdHandler = func(r *photo.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID can not be empty!"})
			return
		}

		intId, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln("Failed to convert the string to an int")
		}

		photo, err := r.GetById(intId)
		if err != nil {
			panic(err)
		}

		if photo == nil {
			c.String(http.StatusNotFound, "")
			return
		}

		representation := representation.PhotoRepresentation{
			ID: photo.ID,
			DateTimeOriginal: photo.DateTimeOriginal,
			Manufacturer: photo.Manufacturer,
			Model: photo.Model,
		}

		c.JSON(http.StatusOK, representation)
	}
}