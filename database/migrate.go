package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

func DBMigrate(db *gorm.DB, entity interface{}) error {
	if err := db.AutoMigrate(entity).Error; err != nil {
		return err
	} else {
		fmt.Println("Auto migrating", reflect.TypeOf(entity).Name(), "...")
	}
	return nil
}