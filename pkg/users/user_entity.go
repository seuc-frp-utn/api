package users

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/seuc-frp-utn/api/pkg/roles"
	"time"
)

type User struct {
	gorm.Model
	UUID string `json:"uuid" gorm:"unique_index"`
	FirstName string `json:"firstName"`
	MiddleName *string `json:"middleName,omitempty"`
	LastName string `json:"lastName"`
	Email string `json:"email" gorm:"unique_index"`
	Birthday time.Time `json:"birthday"`
	Password *string `json:"-"`
	Role roles.Role `json:"role"`
	DNI int64	`json:"dni"`
}

func (u User) Fullname() string {
	if u.MiddleName != nil {
		return fmt.Sprintf("%s %s %s", u.FirstName, *u.MiddleName, u.LastName)
	}
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type UserCreate struct {
	FirstName string `json:"firstName"`
	MiddleName *string `json:"middleName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Birthday time.Time `json:"birthday"`
	Password string `json:"password"`
	DNI int64	`json:"dni"`
}

type UserUpdate struct {
	FirstName *string `json:"firstName"`
	MiddleName *string `json:"middleName"`
	LastName *string `json:"lastName"`
	Email *string `json:"email"`
	Birthday *time.Time `json:"birthday"`
	Password *string `json:"password"`
	Role *roles.Role `json:"role"`
	DNI *uint64	`json:"dni"`
}