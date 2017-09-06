package main

import (
	"github.com/tatthien/go-cms-api/database"
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
		"`author_id` INT (11) UNSIGNED NOT NULL," +
		"`created_at` DATETIME," +
		"`updated_at` DATETIME," +
		"PRIMARY KEY (`id`)," +
		"FOREIGN KEY (`author_id`) REFERENCES users(`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8"
	db.CreateTable("posts", query)
}
