package users

import (
	"errors"
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/auth"
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
		UUID: auth.GenerateUUID(),
		FirstName:  userCreate.FirstName,
		MiddleName: userCreate.MiddleName,
		LastName:   userCreate.LastName,
		Email:      userCreate.Password,
		Birthday:   userCreate.Birthday,
		Password:   hash,
	}

	result, err := (*s.repository).Create(user)
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

func (s Service) Update(uuid string, entity interface{}) (interface{}, error) {
	userUpdate, ok := entity.(UserUpdate)
	if !ok {
		return nil, errors.New("wrong format")
	}

	found, err := s.Get(uuid)
	userFound, ok := found.(User)
	if !ok {
		return nil, errors.New("wrong format")
	}

	user := User{}

	var hash *string
	if userUpdate.Password != nil && !auth.ComparePasswords(*userUpdate.Password, *userFound.Password) {
		if hash, err = auth.GeneratePassword(*userUpdate.Password); err != nil {
			return nil, err
		}
		user.Password = hash
	}

	if err != nil {
		return nil, err
	}

	if userUpdate.FirstName != nil {
		user.FirstName = *userUpdate.FirstName
	}

	if userUpdate.MiddleName != nil {
		user.MiddleName = userUpdate.MiddleName
	}

	if userUpdate.LastName != nil {
		user.LastName = *userUpdate.LastName
	}

	if userUpdate.Email != nil {
		user.Email = *userUpdate.Email
	}

	if userUpdate.Birthday != nil {
		user.Birthday = *userUpdate.Birthday
	}

	result, err := (*s.repository).Update(uuid, user)
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