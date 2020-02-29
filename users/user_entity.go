package users

import (
	"github.com/seuc-frp-utn/api/database"
	"time"
)

type User struct {
	database.Model
	FirstName string `json:"first_name"`
	MiddleName *string `json:"middle_name,omitempty"`
	LastName string `json:"last_name"`
	Email string `json:"email" gorm:"unique"`
	Birthday time.Time `json:"birthday"`
	Password string `json:"-"`
	Salt string `json:"-"`
}

type UserCreate struct {
	FirstName string `json:"first_name"`
	MiddleName *string `json:"middle_name,omitempty"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Birthday time.Time `json:"birthday"`
	Password string `json:"password"`
}