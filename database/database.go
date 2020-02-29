package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"strconv"
)

var (
	Db *gorm.DB
)

// Config represents a set of parameter to configure a database.
type Config struct {
	Host string
	Port int
	User string
	Password string
	Database string
}

func SetupDatabase() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	var port int
	if value, err := strconv.Atoi(os.Getenv("DB_PORT")); err == nil {
		port = value
	} else {
		port = 5432
	}

	config := Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: pass,
		Database: name,
	}

	var err error
	if Db, err = NewPostgresDb(config); err != nil {
		panic("error setting database up")
	}
}

func SetupDatabaseTests() {
	var err error
	path := os.Getenv("DB_TEST_PATH")
	if Db, err = NewSqliteDb(path); err != nil {
		panic("error setting test database up")
	}
}

// New open a new database connection
func NewPostgresDb(config Config) (*gorm.DB, error) {
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

func NewSqliteDb(path string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}
