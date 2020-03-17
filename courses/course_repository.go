package courses

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
	course, ok := entity.(Course)
	if !ok {
		return nil, errors.New("wrong format")
	}
	if pk := r.db.NewRecord(course); !pk {
		return nil, errors.New("course already exists")
	}
	if err := r.db.Create(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r Repository) Read(uuid string) (interface{}, error) {
	var course Course
	if err := r.db.Model(&Course{}).First(&course).Where("uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r Repository) ReadAll() (interface{}, error) {
	var courses []Course
	if err := r.db.Model(&Course{}).Find(&courses).Error; err != nil {
		return nil, err
	}
	return &courses, nil
}

func (r Repository) Update(uuid string, entity interface{}) (interface{}, error) {
	course, ok := entity.(Course)
	if !ok {
		return nil, errors.New("wrong format")
	}
	if err := r.db.Model(&Course{}).Save(&course).Where("uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r Repository) Remove(uuid string) (interface{}, error) {
	found, err := r.Read(uuid)
	if err != nil {
		return nil, err
	}
	course, ok := found.(*Course)
	if !ok {
		return nil, errors.New("wrong format")
	}
	if err := r.db.Delete(course).Error; err != nil {
		return nil, err
	}
	return course, nil
}

func (r Repository) Find(field string, value interface{}) (interface{}, error) {
	var course Course
	query := fmt.Sprintf("%s = ?", field)
	if err := r.db.Model(&Course{}).First(&course).Where(query, value).Error; err != nil {
		return nil, err
	}
	return &course, nil
}
