package users

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/seuc-frp-utn/api/application"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *application.IRepository {
	var r application.IRepository
	r = &Repository{
		db: db,
	}
	return &r
}

func (r Repository) GetDatabase() (*gorm.DB, error) {
	if r.db == nil {
		return nil, errors.New("undefined database")
	}
	return r.db, nil
}

func (r Repository) SetDatabase(db *gorm.DB) error {
	r.db = db
	return nil
}

func (r Repository) Create(entity interface{}) (interface{}, error) {
	panic("implement me")
}

func (r Repository) Read(uuid string) (interface{}, error) {
	panic("implement me")
}

func (r Repository) Update(uuid string, entity interface{}) (interface{}, error) {
	panic("implement me")
}

func (r Repository) Remove(uuid string) (interface{}, error) {
	panic("implement me")
}