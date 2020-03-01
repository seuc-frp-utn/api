package users

import (
	"errors"
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/auth"
	"time"
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

func (s Service) Create(entity interface{}) (interface{}, error) {
	userCreate, ok := entity.(UserCreate)
	if !ok {
		return nil, errors.New("wrong format")
	}

	hash, err := auth.GeneratePassword(userCreate.Password)
	if err != nil {
		return nil, err
	}

	user := User{
		FirstName:  userCreate.FirstName,
		MiddleName: userCreate.MiddleName,
		LastName:   userCreate.LastName,
		Email:      userCreate.Password,
		Birthday:   time.Time{},
		Password:   hash,
	}

	result, err := (*s.repository).Create(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Service) Read(uuid string) (interface{}, error) {
	panic("implement me")
}

func (s Service) ReadAll() (interface{}, error) {
	panic("implement me")
}

func (s Service) Remove(uuid string) (interface{}, error) {
	panic("implement me")
}

func (s Service) Update(uuid string, entity interface{}) (interface{}, error) {
	panic("implement me")
}

func (s Service) Find(field string, value interface{}) (interface{}, error) {
	return (*s.repository).Find(field, value)
}