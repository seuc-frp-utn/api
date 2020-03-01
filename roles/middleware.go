package roles

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(roles Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, ok := c.Get("jwt")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Unassigned role",
			})
			return
		}

		role, ok := value.(Role)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Invalid role",
			})
			return
		}

		if !role.HasRole(roles) {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Not enough permissions",
			})
			return
		}
		c.Next()
	}
}