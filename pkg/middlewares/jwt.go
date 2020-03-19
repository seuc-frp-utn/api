package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/auth"
	"net/http"
)

func JWT(ctx *gin.Context) {
	field := ctx.GetHeader("Authorization")
	if len(field) < 6 {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized - field length"))
		return
	}

	var token string
	fmt.Sscanf(field, "JWT %s", &token)

	if !auth.Sanitize(token) {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid jwt - sanitize failed"))
		return
	}

	jwt, err := auth.Decode(token)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized - error decoding"))
		return
	}
	ctx.Set("jwt", jwt)
}