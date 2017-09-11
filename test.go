package main

import (
	"fmt"
	"log"

	"github.com/tatthien/go-cms-api/database"
	"github.com/tatthien/go-cms-api/model"
)

func main() {
	db := database.Connect()
	defer db.Close()

	// user := model.User{
	// 	Username: "thiennt",
	// 	Password: "111",
	// 	Email:    "x@y.z",
	// }
	// inserted, err := db.InsertUser(user)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// fmt.Printf("%v\n", inserted)
	post := model.Post{
		Title:    "Hello World",
		Content:  "This is an example post.",
		Author:   1,
		PostType: "til",
	}
	post, err := db.InsertPost(post)
	if err != nil {
		log.Fatal(err.Error())
	}

	posts, err := db.GetPosts(0, 10)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%v\n", posts)
}
