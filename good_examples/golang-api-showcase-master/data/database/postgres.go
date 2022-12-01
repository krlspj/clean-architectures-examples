package db

import (
	"fmt"

	"github.com/eldimious/golang-api-showcase/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Import GORM postgres dialect for its side effects, according to GORM docs.
)

// Connect to a database handle from a connection string.
func Connect(configuration *config.Database) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", configuration.Host, configuration.Port, configuration.DB, configuration.User, configuration.Password)
	db, err := gorm.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}
