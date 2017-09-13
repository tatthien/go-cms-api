package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"
	"github.com/tatthien/go-cms-api/database"
	"github.com/tatthien/go-cms-api/model"
)

func main() {
	db := database.Connect()
	defer db.Close()

	// Create `users` table
	query := "CREATE TABLE `users` (" +
		"`id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`username` VARCHAR(20)," +
		"`password` VARCHAR(100)," +
		"`email` VARCHAR(100)," +
		"PRIMARY KEY(`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8"
	db.CreateTable("users", query)

	// Create `posts` table
	query = "CREATE TABLE `posts` (" +
		"`id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`title` TEXT," +
		"`content` TEXT," +
		"`post_type` VARCHAR(20)," +
		"`slug` VARCHAR(255)," +
		"`author_id` INT (11) UNSIGNED NOT NULL," +
		"`created_at` DATETIME," +
		"`updated_at` DATETIME," +
		"PRIMARY KEY (`id`)," +
		"FOREIGN KEY (`author_id`) REFERENCES users(`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8"
	db.CreateTable("posts", query)

	// Insert admin account
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	var user model.User
	user.Username = os.Getenv("ADMIN_USERNAME")
	user.Email = os.Getenv("ADMIN_EMAIL")
	password := os.Getenv("ADMIN_PASSWORD")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	checkError(err)

	user.Password = string(hash)

	admin, err := db.InsertUser(user)
	checkError(err)
	fmt.Println(admin)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
