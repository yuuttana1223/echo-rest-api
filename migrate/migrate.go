package main

import (
	"echo-rest-api/db"
	"echo-rest-api/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
	fmt.Println("Successfully Migrated")
}
