package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to database.")

	return &Database{
		db,
	}
}

// Close the database connection
func (dbfactory Database) Close() {
	dbfactory.db.Close()
}

// CreateTable - create new table
func (dbfactory Database) CreateTable(table, query string) {
	_, err := dbfactory.db.Exec("SET foreign_key_checks = 0")
	checkError(err)

	_, err = dbfactory.db.Exec("DROP TABLE IF EXISTS `" + table + "`")
	checkError(err)

	_, err = dbfactory.db.Exec(query)
	checkError(err)

	fmt.Println("Created table:", table)
}

// Check error
func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
