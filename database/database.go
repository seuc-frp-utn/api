package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Config represents a set of parameter to configure a database.
type Config struct {
	Host string
	Port uint
	User string
	Password string
	Database string
}

// New open a new database connection
func New(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		config.Host,
		config.Port,
		config.User,
		config.Database,
		config.Password,
	)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
