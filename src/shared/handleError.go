package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err.Error(),
			"status": false,
		})

		return
	}

}
