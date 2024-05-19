package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Cache *cache.Cache
)

type Person struct {
	ID   int `gorm:"primarykey,autoincrement:true"`
	Name string
	Age  int
}

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Person{})
	DB = db

	Cache = cache.New(5*time.Minute, 10*time.Minute)
}

func GetPerson(ID int) (*Person, error) {
	if x, found := Cache.Get(fmt.Sprintf("person_%d", ID)); found {
		fmt.Println("got from cache")
		return x.(*Person), nil
	}

	person := Person{}
	if err := DB.First(&person, ID).Error; err != nil {
		return nil, err
	}
	Cache.Set(fmt.Sprintf("person_%d", ID), &person, cache.DefaultExpiration)
	fmt.Println("got from db")
	return &person, nil
}

func main() {
	person := Person{Name: "Jinzhu", Age: 18}
	DB.Create(&person)

	gotPerson, _ := GetPerson(person.ID)
	fmt.Println(gotPerson.Name)
	fmt.Println(gotPerson.Age)

	gotPerson, _ = GetPerson(person.ID)
	fmt.Println(gotPerson.Name)
	fmt.Println(gotPerson.Age)
}
