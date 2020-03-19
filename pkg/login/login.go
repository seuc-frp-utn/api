package login

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/pkg/auth"
	"github.com/seuc-frp-utn/api/pkg/users"
	"net/http"
)

func Handler(ctx *gin.Context) {
	var input CredentialsInput
	ctx.BindJSON(&input)
	
	found, err := (*users.UserService).Find("email", input.Email)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	user, ok := found.(*users.User)
	if !ok {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid user"))
		return
	}
	
	if !auth.ComparePasswords(input.Password, *user.Password) {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid credentials"))
		return
	}
	
	jwt := auth.JWT{
		UUID:  user.UUID,
		Name:  user.Fullname(),
		Email: user.Email,
		Roles: user.Role,
	}
	
	token, expiresAt, err := auth.Encode(jwt)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("invalid credentials"))
	}
	
	credentials := CredentialsOutput{
		Token:     *token,
		ExpiresAt: *expiresAt,
		Subject: user.UUID,
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"access_token": credentials.Token,
		"sub": credentials.Subject,
		"exp": credentials.ExpiresAt,
	})
}
