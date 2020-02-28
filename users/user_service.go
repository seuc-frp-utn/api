package users

import (
	"github.com/seuc-frp-utn/api/application"
)

type Service struct {
	repository *application.IRepository
}

func NewService(repository *application.IRepository) *application.IService {
	var s application.IService
	s = &Service{
		repository: repository,
	}
	return &s
}

func (s Service) GetRepository() (*application.IRepository, error) {
	panic("implement me")
}

func (s Service) SetRepository(repository *application.IRepository) error {
	panic("implement me")
}

func (s Service) Create(entity interface{}) (interface{}, error) {
	panic("implement me")
}

func (s Service) Read(uuid string) (interface{}, error) {
	panic("implement me")
}

func (s Service) Remove(uuid string) (interface{}, error) {
	panic("implement me")
}

func (s Service) Update(uuid string, entity interface{}) (interface{}, error) {
	panic("implement me")
}