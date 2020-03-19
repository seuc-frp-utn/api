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

type IServiceCreate interface {
	Create(entity reflect.Value) (interface{}, error)
}

type IServiceGet interface {
	Get(uuid string) (interface{}, error)
}

type IServiceGetAll interface {
	GetAll() (interface{}, error)
}

type IServiceUpdate interface {
	Update(uuid string, entity reflect.Value) (interface{}, error)
}

type IServiceRemove interface {
	Remove(uuid string) (interface{}, error)
}

type IServiceFind interface {
	Find(field string, value interface{}) (interface{}, error)
}