package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/auth"
	"net/http"
)

func UUID(c *gin.Context) {
	uuid := c.Param("uuid")
	if !auth.IsUUID(uuid) {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid UUID",
			},
		)
		return
	}
}