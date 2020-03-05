package users

import (
	"github.com/jinzhu/gorm"
	"github.com/seuc-frp-utn/api/roles"
	"time"
)

type User struct {
	gorm.Model
	UUID string `json:"uuid" gorm:"unique_index"`
	FirstName string `json:"first_name"`
	MiddleName *string `json:"middle_name,omitempty"`
	LastName string `json:"last_name"`
	Email string `json:"email" gorm:"unique_index"`
	Birthday time.Time `json:"birthday"`
	Password *string `json:"-,omitempty"`
	Role roles.Role `json:"role"`
}

type UserCreate struct {
	FirstName string `json:"first_name"`
	MiddleName *string `json:"middle_name,omitempty"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Birthday time.Time `json:"birthday"`
	Password string `json:"password"`
}

type UserUpdate struct {
	FirstName *string `json:"first_name,omitempty"`
	MiddleName *string `json:"middle_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Email *string `json:"email,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
	Password *string `json:"password,omitempty"`
	Role *roles.Role `json:"role,omitempty"`
}