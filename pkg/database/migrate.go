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