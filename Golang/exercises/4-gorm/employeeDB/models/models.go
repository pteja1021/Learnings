package models

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
	Department	Department	
}

type Department struct {
	gorm.Model
	Name      string
	Location  string
	Employees []Employee
}

func InitiateDB(){

	//reading command line flags 
	dbName := flag.String("db","gorm","a string")
	
	flag.Parse()

	//establishing connection	
	db, err := gorm.Open("postgres", fmt.Sprintf("user=postgres password=bcdev123 dbname=%v sslmode=disable",*dbName))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

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
	db.Save(&manager)

	newEmployee := Employee{
		Name:         "naga",
		Role:         "developer",
		Manager:      &manager,
		DepartmentID: department.ID,
	}	
	db.Save(&newEmployee)

	// Finding testing dept and inserting an employee in it with manager as vamshee
	testingDepartment := Department{}
	db.Where(&Department{Name: "Testing"}).Find(&testingDepartment)
	db.Save(&Employee{Name: "Smith",Role: "tester", Manager: &manager, Department: testingDepartment})
	

	allEmployees := []Employee{}
	db.Preload("Manager").Preload("Department").Find(&allEmployees)

	// Printing all employees
	for index := range allEmployees{
		if allEmployees[index].Manager==nil{
			fmt.Println(allEmployees[index].Name,"No manager",allEmployees[index].Department.Name)
		} else {
			fmt.Println(allEmployees[index].Name,allEmployees[index].Manager.Name,allEmployees[index].Department.Name)
		}
	}
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
