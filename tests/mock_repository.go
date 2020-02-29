package tests

import (
	"github.com/jinzhu/gorm"
)

type MockRepository struct {
	GetDatabaseMock func() (*gorm.DB, error)
	SetDatabaseMock func(db *gorm.DB) error
	CreateMock func(entity interface{}) (interface{}, error)
	ReadMock func(uuid string) (interface{}, error)
	ReadAllMock func() (interface{}, error)
	UpdateMock func(uuid string, entity interface{}) (interface{}, error)
	RemoveMock func(uuid string) (interface{}, error)
}

func (s MockRepository) GetDatabase() (*gorm.DB, error) {
	return s.GetDatabaseMock()
}

func (s MockRepository) SetDatabase(db *gorm.DB) error {
	return s.SetDatabaseMock(db)
}

func (s MockRepository) Create(entity interface{}) (interface{}, error) {
	return s.CreateMock(entity)
}

func (s MockRepository) Read(uuid string) (interface{}, error) {
	return s.ReadMock(uuid)
}

func (s MockRepository) ReadAll() (interface{}, error) {
	return s.ReadAllMock()
}

func (s MockRepository) Remove(uuid string) (interface{}, error) {
	return s.RemoveMock(uuid)
}

func (s MockRepository) Update(uuid string, entity interface{}) (interface{}, error) {
	return s.UpdateMock(uuid, entity)
}
