package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/pkg/application"
	"github.com/seuc-frp-utn/api/pkg/tests"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)


func TestController_GetService(t *testing.T) {

}

func TestController_SetService(t *testing.T) {

}

func TestController_Create(t *testing.T) {
	r := gin.Default()

	root := r.Group("/")
	{
		RegisterDirectTest(root)
	}

	var mock application.IService

	mock = tests.MockService{
		CreateMock: func(entity reflect.Value) (interface{}, error) {
				user, ok := entity.Interface().(UserCreate)
				if !ok {
					return nil, errors.New("wrong format")
				}
				return User{
					FirstName:  user.FirstName,
					MiddleName: user.MiddleName,
					LastName:   user.LastName,
					Email:      user.Email,
					Birthday:   time.Time{},
					Password:   nil,
				}, nil
		},
	}

	(*UserController).SetService(&mock)
	
	userCreate := UserCreate{
		FirstName:  "Test",
		MiddleName: nil,
		LastName:   "Test",
		Email:      "test@test.org",
		Birthday:   time.Time{},
		Password:   "1234",
	}

	mjson, err := json.Marshal(userCreate)
	if err != nil {
		t.Fail()
	}
	body := bytes.NewReader(mjson)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		t.Fail()
	}

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var result User
	json.Unmarshal(w.Body.Bytes(), &result)

	assert.Equal(t, userCreate.FirstName, result.FirstName)
	assert.Equal(t, userCreate.LastName, result.LastName)
	assert.Equal(t, userCreate.Email, result.Email)
	assert.Equal(t, userCreate.Birthday, result.Birthday)
	assert.Nil(t, result.Password)
}