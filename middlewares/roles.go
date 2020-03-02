package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/auth"
	"github.com/seuc-frp-utn/api/roles"
	"net/http"
)

func Roles(roles roles.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, ok := c.Get("jwt")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Invalid JWT",
			})
			return
		}

		jwt, ok := value.(auth.JWT)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Invalid JWT",
			})
			return
		}

		uuid := c.Param("uuid")
		if jwt.Roles.IsUser() && jwt.UUID != uuid {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Not enough permissions",
			})
		}

		if !jwt.Roles.HasRole(roles) {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Not enough permissions",
			})
			return
		}
		c.Next()
	}
}