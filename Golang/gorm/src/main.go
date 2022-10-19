package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct  { // defining struct to create a db table with same fields as the struct
	gorm.Model // composing with gorm model, will directly give access to default fields, id,created_at,updated_at,deleted_at
	Username string
	FirstName string // see how these names are reflected in actual db. gorm converts pascal casing to underscores 
	LastName string 
}

func main(){
	// to establish a connection to DB from gorm
	db,err := gorm.Open("postgres","user=postgres password=bcdev123 dbname=gorm sslmode=disable") // if db doesnt exists it throws an error
	if err!=nil {
		panic(err.Error())
	}
	defer db.Close()

	dbase := db.DB() // establishes the database connection directly
	defer dbase.Close()

	err = dbase.Ping() // Pinging to check if the db connection is establised or not
	if err!=nil {
		panic(err.Error())
	}
	fmt.Println("connection established")

	db.DropTableIfExists(&User{})
	db.CreateTable(&User{}) // creating a table

	// CRUD OPERATIONS

	// CREATE
	user1 := User { // creating a user for "users" table, the id field will be automatically given by gorm since it is PK and integer- autoinc
		Username: "teja_p",
		FirstName: "teja",
		LastName: "p",
	}
	// to check if a record has been inserted or not
	fmt.Println(db.NewRecord(&user1)) // returns a bool- false if inserted, true if not inserted
	db.Create(&user1) // inserting a record into User associated table
	db.Create(&User{Username: "naga",FirstName: "Naga", LastName: "Reddy"})

	// READ
	firstUser := User{}
	db.First(&firstUser) // getting first record from users table
	fmt.Println("first user",firstUser)
	
	lastUser := User{}
	db.Last(&lastUser) // getting last record from users table
	fmt.Println("last user",lastUser)

	// UPDATE

	findUser := User{Username: "teja_p"}
	db.Where(&findUser).First(&findUser) // find in users where findUser conditions are met,
	// grab the first such record(Where) and put that fisrt record into findUser()

	findUser.LastName="puvvula" // update the properties
	
	db.Save(&findUser) // persist the changes made to the found object

	updatedUser := User{}
	db.First(&updatedUser) // check if the record has been updated or not
	fmt.Println("updated first user is ",updatedUser) 

	// DELETE	
	db.Where(&User{Username: "teja_p"}).Delete(&User{})
}