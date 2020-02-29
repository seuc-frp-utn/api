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
	user, ok := entity.(User)
	if !ok {
		return nil, errors.New("wrong format")
	}
	if pk := r.db.NewRecord(user); pk {
		return nil, errors.New("user already exists")
	}
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r Repository) Read(uuid string) (interface{}, error) {
	var result User
	if err := r.db.Model(&User{}).First(&result).Where("uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
}

func (r Repository) ReadAll() (interface{}, error) {
	r.db.Find(&User{})
}

func (r Repository) Update(uuid string, entity interface{}) (interface{}, error) {
	panic("implement me")
}

func (r Repository) Remove(uuid string) (interface{}, error) {
	panic("implement me")
}