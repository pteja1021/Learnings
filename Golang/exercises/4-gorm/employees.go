package main

import (
	"fmt"
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Employee struct {
	gorm.Model
	Name         string
	Role         string
	ManagerID    *uint
	Manager      *Employee	
	DepartmentID uint
}

type Department struct {
	gorm.Model
	Name      string
	Location  string
	Employees []Employee
}

func main() {

	//reading command line flags 
	dbName := flag.String("db","gorm","a string")
	
	flag.Parse()

	//establishing connection	
	db, err := gorm.Open("postgres", fmt.Sprintf("user=postgres password=bcdev123 dbname=%v sslmode=disable",*dbName))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("connection established")

	//Creating Tables
	db.DropTableIfExists(&Employee{}, &Department{})
	db.CreateTable(&Employee{}, &Department{})

	//Populating Department table
	for _, value := range DepartmentData() {
		db.Create(&value)
	}

	department := Department{
		Name:      "Infrastructure",
		Location:  "Hyderabad",
		Employees: []Employee{},
	}
	db.Save(&department)

	manager := Employee{
		Name: "vamshee",
		Role: "manager",
		DepartmentID: department.ID,
	}
	
	newEmployee := Employee{
		Name:         "naga",
		Role:         "developer",
		Manager:      &manager,
		DepartmentID: department.ID,
	}
	
	db.Debug().Save(&manager)
	db.Debug().Save(&newEmployee)
	
	new_department := Department{}
	db.Preload("Employees").Where(&department).Find(&new_department)
	fmt.Println(new_department.Employees)
}


func DepartmentData() []Department {
	departments := []Department{
		{Name: "Development", Location: "Hyderabad"},
		{Name: "Testing", Location: "Bangalore"},
		{Name: "Operations", Location: "Pune"},
		{Name: "Deployment", Location: "Chennai"},
	}
	return departments
}
