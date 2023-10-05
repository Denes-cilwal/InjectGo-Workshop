package infrastructure

import (
	"InjectGo-Workshop/lib"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Database modal
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(env *lib.Env) (*Database, error) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort)
	log.Println("opening db connection")

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Println("Url: ", url)
		log.Panic(err)
	}

	// Set the default database for this DB instance
	err = db.Exec("USE " + env.DBName).Error
	if err != nil {
		log.Panic(err)
	}

	// Automatically create the database if it doesn't exist
	err = db.Exec("CREATE DATABASE IF NOT EXISTS " + env.DBName).Error
	if err != nil {
		fmt.Println("tet")
		return nil, err
	}
	database := &Database{DB: db}
	return database, nil
}
