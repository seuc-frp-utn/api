package database

import "github.com/jinzhu/gorm"

type Model struct {
	gorm.Model
	UUID string `json:"uuid"`
}