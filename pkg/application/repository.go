package application

import "github.com/jinzhu/gorm"

type IRepository interface {
	GetDatabase() (*gorm.DB, error)
	SetDatabase(db *gorm.DB) error
	Create(entity interface{}) (interface{}, error)
	Get(uuid string) (interface{}, error)
	GetAll() (interface{}, error)
	Update(uuid string, entity interface{}) (interface{}, error)
	Remove(uuid string) (interface{}, error)
	Find(field string, value interface{}) (interface{}, error)
}