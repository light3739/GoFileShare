package configs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectToDatabase() (*gorm.DB, error) {
	host := DbHost
	port := DbPort // PostgreSQL default port is 5432
	user := DbUser
	password := DbPassword
	dbname := DbName

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	// Open a connection to the database
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database!")
	return db, nil
}
