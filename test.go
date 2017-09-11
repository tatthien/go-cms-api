package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/tatthien/go-cms-api/database"
	"github.com/tatthien/go-cms-api/model"
)

func main() {
	db := database.Connect()
	defer db.Close()

	password, _ := bcrypt.GenerateFromPassword([]byte("111111"), bcrypt.DefaultCost)
	user := model.User{
		Username: "thiennt01",
		Password: string(password),
		Email:    "tatthien.contact@gmail.com",
	}
	inserted, err := db.InsertUser(user)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(inserted)
}
