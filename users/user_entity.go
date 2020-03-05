package users

import (
	"fmt"
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
	Password *string `json:"-"`
	Role roles.Role `json:"role"`
}

func (u User) Fullname() string {
	if u.MiddleName != nil {
		return fmt.Sprintf("%s %s %s", u.FirstName, *u.MiddleName, u.LastName)
	}
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
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