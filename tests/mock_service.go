package tests

import "github.com/seuc-frp-utn/api/application"

type MockService struct {
	GetRepositoryMock func() (*application.IRepository, error)
	SetRepositoryMock func(repository *application.IRepository) error
	CreateMock func(entity interface{}) (interface{}, error)
	ReadMock func(uuid string) (interface{}, error)
	ReadAllMock func() (interface{}, error)
	UpdateMock func(uuid string, entity interface{}) (interface{}, error)
	RemoveMock func(uuid string) (interface{}, error)
	FindMock func(field string, value interface{}) (interface{}, error)
}

func (s MockService) GetRepository() (*application.IRepository, error) {
	return s.GetRepositoryMock()
}

func (s MockService) SetRepository(repository *application.IRepository) error {
	return s.SetRepositoryMock(repository)
}

func (s MockService) Create(entity interface{}) (interface{}, error) {
	return s.CreateMock(entity)
}

func (s MockService) Read(uuid string) (interface{}, error) {
	return s.ReadMock(uuid)
}

func (s MockService) ReadAll() (interface{}, error) {
	return s.ReadAllMock()
}

func (s MockService) Remove(uuid string) (interface{}, error) {
	return s.RemoveMock(uuid)
}

func (s MockService) Update(uuid string, entity interface{}) (interface{}, error) {
	return s.UpdateMock(uuid, entity)
}

func (s MockService) Find(field string, value interface{}) (interface{}, error) {
	return s.FindMock(field, value)
}