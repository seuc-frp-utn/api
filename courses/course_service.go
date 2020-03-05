package courses

import (
	"errors"
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/auth"
	"github.com/seuc-frp-utn/api/users"
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
	courseCreate, ok := entity.(CourseCreate)
	if !ok {
		return nil, errors.New("wrong format")
	}

	found, err := (*users.UserService).Get(courseCreate.Teacher)
	if err != nil {
		return nil, errors.New("teacher does not exist")
	}
	teacher, ok := found.(users.User)
	if !ok {
		return nil, errors.New("invalid teacher")
	}
	if teacher.Role.IsTeacher() == false {
		return nil, errors.New("selected user is not a teacher")
	}

	course := Course{
		UUID: auth.GenerateUUID(),
		Name:        courseCreate.Name,
		Description: courseCreate.Description,
		Image:       courseCreate.Image,
		Started:     courseCreate.Started,
		Finished:    courseCreate.Finished,
		Hours:       courseCreate.Hours,
		Price:       courseCreate.Price,
		Months:      courseCreate.Months,
		Link:        courseCreate.Link,
		Teacher:	courseCreate.Teacher,
	}

	result, err := (*s.repository).Create(course)
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
	courseUpdate, ok := entity.(CourseUpdate)
	if !ok {
		return nil, errors.New("wrong format")
	}

	_, err := s.Get(uuid)
	if err != nil {
		return nil, errors.New("course does not exist")
	}

	course := Course{}

	if courseUpdate.Name != nil {
		course.Name = *courseUpdate.Name
	}
	if courseUpdate.Description != nil {
		course.Description = *courseUpdate.Description
	}
	if courseUpdate.Image != nil {
		course.Image = *courseUpdate.Image
	}
	if courseUpdate.Started != nil {
		course.Started = *courseUpdate.Started
	}
	if courseUpdate.Finished != nil {
		course.Finished = *courseUpdate.Finished
	}
	if courseUpdate.Link != nil {
		course.Link = *courseUpdate.Link
	}
	if courseUpdate.Hours != nil {
		course.Hours = *courseUpdate.Hours
	}
	if courseUpdate.Price != nil {
		course.Price = *courseUpdate.Price
	}
	if courseUpdate.Months != nil {
		course.Months = *courseUpdate.Months
	}
	if courseUpdate.Teacher != nil {
		if _, err := (*users.UserService).Get(*courseUpdate.Teacher); err == nil {
			course.Teacher = *courseUpdate.Teacher
		}
	}

	result, err := (*s.repository).Update(uuid, course)
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