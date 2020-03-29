package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	if Db, err = NewMySQLDb(config); err != nil {
		panic(fmt.Sprintf("[FATAL] Error setting up database: %v", err))
	}
	Db.LogMode(false)
}

func SetupDatabaseTests() {
	var err error
	path := os.Getenv("DB_TEST_PATH")
	if Db, err = NewSqliteDb(path); err != nil {
		panic(fmt.Sprintf("[FATAL] Error setting up database: %v", err))
	}
	Db.LogMode(true)
}

// New open a new database connection
func NewPostgresDb(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
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

func NewMySQLDb(config Config) (*gorm.DB, error) {
	// user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.User, config.Password, config.Host, config.Port, config.Database)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db,nil
}

func NewSqliteDb(path string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Close() {
	Db.Close()
}