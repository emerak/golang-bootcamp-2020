package main

import (
	"github.com/emerak/golang-bootcamp-2020/config"
	"github.com/emerak/golang-bootcamp-2020/domain/model"
	"github.com/emerak/golang-bootcamp-2020/infrastructure/datastore"
)

func main() {
	config.ReadConfig()
	db := datastore.NewDB()
	db.LogMode(true)

	user1 := model.User{FirstName: "John", LastName: "Doe", Username: "johndoe"}
	db.Create(&user1)

	user2 := model.User{FirstName: "Jane", LastName: "Dee", Username: "janedee"}
	db.Create(&user2)

	defer db.Close()
}
