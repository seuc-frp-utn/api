package database

import (
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB, entity interface{}) error {
	if err := db.AutoMigrate(entity).Error; err != nil {
		return err
	}
	return nil
}

func AddDefaultData(db *gorm.DB, model interface{}, data interface{}) error {
	err := db.Model(&model).Save(&data).Error
	if err != nil {
		return err
	}
	return nil
}