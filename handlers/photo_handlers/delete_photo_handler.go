package photo_handlers

import (
	"UploadService/domain/photo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var DeleteUploadHandler = func(r *photo.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID can not be empty!"})
			return
		}

		intId, err := strconv.Atoi(id)

		err = r.Delete(intId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, "Deleted.")
	}
}