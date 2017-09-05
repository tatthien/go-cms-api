package database

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

// Connect to database
func Connect() *Database {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	dbName := os.Getenv("DATABASE_NAME")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")

	db, err := sql.Open("mysql", dbUsername + ":" + dbPassword + "@/" + dbName)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to database.")

	return &Database{
		db
	}
}

// Close the database connection
func (dbfactory Database) Close() {
	dbfactory.db.Close()
}

func (dbfactory Database) CreateTable(table, query string) error {
	
}