package diplomas

import (
	"errors"
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/auth"
	"github.com/seuc-frp-utn/api/courses"
	"github.com/seuc-frp-utn/api/users"
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

	if !auth.IsUUID(diplomaCreate.Course) {
		return nil, errors.New("invalid course")
	}

	if !auth.IsUUID(diplomaCreate.Dean) {
		return nil, errors.New("invalid dean")
	}

	if !auth.IsUUID(diplomaCreate.Teacher) {
		return nil, errors.New("invalid teacher")
	}

	if !auth.IsUUID(diplomaCreate.Secretary) {
		return nil, errors.New("invalid secretary")
	}

	var err error

	_, err = (*courses.CourseService).Get(diplomaCreate.Course)
	if err != nil {
		return nil, err
	}

	_, err = (*users.UserService).Get(diplomaCreate.Teacher)
	if err != nil {
		return nil, err
	}

	_, err = (*users.UserService).Get(diplomaCreate.Secretary)
	if err != nil {
		return nil, err
	}

	_, err = (*users.UserService).Get(diplomaCreate.Dean)
	if err != nil {
		return nil, err
	}


	token, err := auth.GenerateNanoUUID()
	if err != nil {
		return nil, err
	}

	diploma := Diploma{
		UUID:      auth.GenerateUUID(),
		Token:     token,
		Course:    diplomaCreate.Course,
		Dean:      diplomaCreate.Dean,
		Secretary: diplomaCreate.Secretary,
		Teacher:   diplomaCreate.Teacher,
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
	diplomaUpdate, ok := entity.Interface().(DiplomaUpdate)
	if !ok {
		return nil, errors.New("wrong format")
	}

	found, err := s.Get(uuid)
	if err != nil {
		return nil, errors.New("diploma does not exist")
	}
	diploma, ok := found.(Diploma)

	if diplomaUpdate.Course != nil {
		_, err = (*courses.CourseService).Get(*diplomaUpdate.Course)
		if err != nil {
			return nil, err
		}
		diploma.Course = *diplomaUpdate.Course
	}

	if diplomaUpdate.Teacher != nil {
		_, err = (*users.UserService).Get(*diplomaUpdate.Teacher)
		if err != nil {
			return nil, err
		}
		diploma.Teacher = *diplomaUpdate.Teacher
	}

	if diplomaUpdate.Secretary != nil {
		_, err = (*users.UserService).Get(*diplomaUpdate.Secretary)
		if err != nil {
			return nil, err
		}
		diploma.Secretary = *diplomaUpdate.Secretary
	}

	if diplomaUpdate.Dean != nil {
		_, err = (*users.UserService).Get(*diplomaUpdate.Dean)
		if err != nil {
			return nil, err
		}
		diploma.Dean = *diplomaUpdate.Dean
	}


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