package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/auth"
	"net/http"
)

func JWT(ctx *gin.Context) {
	field := ctx.GetHeader("Authorization")
	if len(field) < 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var token string
	fmt.Sscanf(field, "JWT %s", &token)

	if !auth.Sanitize(token) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid JWT",
		})
		return
	}

	jwt, err := auth.Decode(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	ctx.Set("jwt", jwt)
	ctx.Next()
}