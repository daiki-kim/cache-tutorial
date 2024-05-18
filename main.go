package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Person struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Person{})

	person := Person{Name: "Jinzhu", Age: 18}
	db.Create(&person)

	var result Person
	db.First(&result, person.ID)
	fmt.Println(result)
}
