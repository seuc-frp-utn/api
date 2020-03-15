package diplomas

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
	diploma, ok := entity.(Diploma)
	if !ok {
		return nil, errors.New("wrong format")
	}
	if pk := r.db.NewRecord(diploma); !pk {
		return nil, errors.New("diploma already exists")
	}
	if err := r.db.Create(&diploma).Error; err != nil {
		return nil, err
	}
	return &diploma, nil
}

func (r Repository) Read(uuid string) (interface{}, error) {
	var diploma Diploma
	if err := r.db.Model(&Diploma{}).First(&diploma).Where("uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &diploma, nil
}

func (r Repository) ReadAll() (interface{}, error) {
	var diplomas []Diploma
	if err := r.db.Model(&Diploma{}).Find(&diplomas).Error; err != nil {
		return nil, err
	}
	return &diplomas, nil
}

func (r Repository) Update(uuid string, entity interface{}) (interface{}, error) {
	diploma, ok := entity.(Diploma)
	if !ok {
		return nil, errors.New("wrong format")
	}
	if err := r.db.Model(&Diploma{}).Save(&diploma).Where("uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &diploma, nil
}

func (r Repository) Remove(uuid string) (interface{}, error) {
	var diploma Diploma
	if err := r.db.Model(&Diploma{}).Where("uuid = ?").Delete(&diploma).Error; err != nil {
		return nil, err
	}
	return &diploma, nil
}

func (r Repository) Find(field string, value interface{}) (interface{}, error) {
	var diploma Diploma
	query := fmt.Sprintf("%s = ?", field)
	if err := r.db.Model(&Diploma{}).First(&diploma).Where(query, value).Error; err != nil {
		return nil, err
	}
	return &diploma, nil
}