package users

import (
	"errors"
	"fmt"
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
	var user User
	if err := r.db.Model(&User{}).First(&user).Where("uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r Repository) ReadAll() (interface{}, error) {
	var users []User
	if err := r.db.Model(&[]User{}).Find(users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r Repository) Update(uuid string, entity interface{}) (interface{}, error) {
	user, ok := entity.(User)
	if !ok {
		return nil, errors.New("wrong format")
	}
	if err := r.db.Model(&User{}).Save(&user).Where("uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r Repository) Remove(uuid string) (interface{}, error) {
	var user User
	if err := r.db.Model(&User{}).Where("uuid = ?").Delete(user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r Repository) Find(field string, value interface{}) (interface{}, error) {
	var user User
	query := fmt.Sprintf("%s = ?", field)
	if err := r.db.Model(&User{}).First(user).Where(query, value).Error; err != nil {
		return nil, err
	}
}