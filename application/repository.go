package application

import "github.com/jinzhu/gorm"

type IRepository interface {
	GetDatabase() (*gorm.DB, error)
	SetDatabase(db *gorm.DB) error
	Create(entity interface{}) (interface{}, error)
	Read(uuid string) (interface{}, error)
	ReadAll() (interface{}, error)
	Update(uuid string, entity interface{}) (interface{}, error)
	Remove(uuid string) (interface{}, error)
}