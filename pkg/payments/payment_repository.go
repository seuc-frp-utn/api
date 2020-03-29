package payments

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/seuc-frp-utn/api/pkg/application"
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
	payment, ok := entity.(Payment)
	if !ok {
		return nil, errors.New("wrong format")
	}
	if pk := r.db.NewRecord(payment); !pk {
		return nil, errors.New("payment already exists")
	}
	if err := r.db.Create(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r Repository) Get(uuid string) (interface{}, error) {
	var payment Payment
	if err := r.db.Model(&Payment{}).First(&payment).Where("uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r Repository) GetAll() (interface{}, error) {
	var payments []Payment
	if err := r.db.Model(&Payment{}).Find(&payments).Error; err != nil {
		return nil, err
	}
	return &payments, nil
}

func (r Repository) Update(uuid string, entity interface{}) (interface{}, error) {
	payment, ok := entity.(Payment)
	if !ok {
		return nil, errors.New("wrong format")
	}
	if err := r.db.Model(&Payment{}).Save(&payment).Where("uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r Repository) Remove(uuid string) (interface{}, error) {
	var payment Payment
	err := r.db.Model(&Payment{}).Where("uuid = ?", uuid).First(&payment).Error
	if err != nil {
		return nil, err
	}
	if err := r.db.Delete(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r Repository) Find(field string, value interface{}) (interface{}, error) {
	var payment Payment
	query := fmt.Sprintf("%s = ?", field)
	if err := r.db.Model(&Payment{}).Where(query, value).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}
