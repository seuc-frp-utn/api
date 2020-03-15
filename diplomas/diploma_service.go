package diplomas

import (
	"errors"
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/auth"
	"reflect"
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
	if s.repository == nil {
		return nil, errors.New("undefined repository")
	}
	return s.repository, nil
}

func (s Service) SetRepository(repository *application.IRepository) error {
	s.repository = repository
	return nil
}

func (s Service) Create(entity reflect.Value) (interface{}, error) {
	diplomaCreate, ok := entity.Interface().(DiplomaCreate)
	if !ok {
		return nil, errors.New("wrong format")
	}

	diploma := Diploma{
		UUID: auth.GenerateUUID(),
	}

	result, err := (*s.repository).Create(diploma)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Service) Get(uuid string) (interface{}, error) {
	result, err := (*s.repository).Read(uuid)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Service) GetAll() (interface{}, error) {
	result, err := (*s.repository).ReadAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Service) Remove(uuid string) (interface{}, error) {
	result, err := (*s.repository).Remove(uuid)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Service) Update(uuid string, entity reflect.Value) (interface{}, error) {
	diplomaUpdate, ok := reflect.ValueOf(entity).Interface().(DiplomaUpdate)
	if !ok {
		return nil, errors.New("wrong format")
	}

	_, err := s.Get(uuid)
	if err != nil {
		return nil, errors.New("diploma does not exist")
	}

	diploma := Diploma{}

	result, err := (*s.repository).Update(uuid, diploma)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Service) Find(field string, value interface{}) (interface{}, error) {
	result, err := (*s.repository).Find(field, value)
	if err != nil {
		return nil, err
	}
	return result, nil
}