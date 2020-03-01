package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(ctx *gin.Context) {
	field := ctx.GetHeader("Authorization")
	if len(field) < 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var token string
	fmt.Sscanf(field, "JWT %s", &token)

	if !Sanitize(token) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid JWT",
		})
		return
	}

	jwt, err := Decode(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	ctx.Set("jwt", jwt)
	ctx.Next()
}