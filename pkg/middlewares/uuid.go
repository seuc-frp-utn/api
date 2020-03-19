package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/auth"
	"net/http"
)

func UUID(c *gin.Context) {
	uuid := c.Param("uuid")
	if !auth.IsUUID(uuid) {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid UUID"))
		return
	}
}