package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/pkg/auth"
	"github.com/seuc-frp-utn/api/pkg/roles"
	"net/http"
)

func Roles(roles roles.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, ok := c.Get("jwt")
		if !ok {
			c.AbortWithError(http.StatusForbidden, errors.New("invalid jwt - roles"))
			return
		}

		jwt, ok := value.(*auth.JWT)
		if !ok {
			c.AbortWithError(http.StatusForbidden, errors.New("invalid jwt - roles casting"))
			return
		}

		uuid := c.Param("uuid")
		if len(uuid) > 0 {
			if jwt.Roles.IsUser() && !jwt.Roles.IsAdmin() && !jwt.Roles.IsOperator() && jwt.UUID != uuid {
				c.AbortWithError(http.StatusForbidden, errors.New("not enough permissions - uuid mismatch"))
				return
			}
		}

		if !jwt.Roles.HasRole(roles) {
			c.AbortWithError(http.StatusForbidden, errors.New("not enough permissions"))
			return
		}
	}
}