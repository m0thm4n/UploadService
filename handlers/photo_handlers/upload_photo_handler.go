package photo_handlers

import (
	"github.com/m0thm4n/UploadService/domain/photo"
	"github.com/m0thm4n/UploadService/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"os"
)

var UploadPhotoHandler = func(r *photo.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		filePath := "/home/timmy/photos/"

		f, uploadedFile, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error": true,
			})

			return
		}

		fileName := uploadedFile.Filename

		out, err := os.Create(filePath + fileName)

		defer f.Close()

		if _, err := io.Copy(out, f); err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"error": true,
			})

			return
		}

		u, err := url.Parse(filePath + uploadedFile.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"Error": true,
			})

			return
		}

		values := utils.ExtractPhoto(filePath + uploadedFile.Filename)
		dateTimeOrignial := values["DateTimeOriginal"]
		manufacturer := values["Manufacturer"]
		model := values["Model"]

		photo := &photo.Photo{
			DateTimeOriginal: dateTimeOrignial,
			Manufacturer: manufacturer,
			Model: model,
		}

		err = r.Insert(photo)
		if err != nil {
			panic(err)
		}


		c.JSON(http.StatusOK, gin.H{
			"message": "File Uploaded to the server successfully!",
			"pathname": u.EscapedPath(),
		})
	}
}
