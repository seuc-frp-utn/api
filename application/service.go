package application

import "reflect"

type IService interface {
	GetRepository() (*IRepository, error)
	SetRepository(repository *IRepository) error
	Create(entity reflect.Value) (interface{}, error)
	Get(uuid string) (interface{}, error)
	GetAll() (interface{}, error)
	Update(uuid string, entity reflect.Value) (interface{}, error)
	Remove(uuid string) (interface{}, error)
	Find(field string, value interface{}) (interface{}, error)
}