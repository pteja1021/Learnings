package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/employeeDB/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main(){
	conn, err:= grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()),grpc.WithBlock()) // with block will block until conn is made
	if err!=nil{
		log.Fatal("Did not connect", err.Error())
	}
	defer conn.Close()
	c:= pb.NewEmployeeDatabaseCrudClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	created_employee,err := c.AddEmployee(ctx, &pb.NewEmployee{Name: "Will", DepartmentId: 2, Role: "Tester"})
	if err!=nil{
		log.Fatal(err.Error())
	}
	log.Printf("id is %v, name is %v",created_employee.GetId(),created_employee.GetName())

	allEmployees,err := c.GetEmployees(ctx,&pb.EmptyEmployeeRequest{})
	if err!=nil{
		log.Printf("error getting all employees")
	}
	for _,employee := range allEmployees.Employees{
		fmt.Println(employee.GetName(),employee.GetId())
	}

	updated_employee,err := c.UpdateRoleOfEmployee(ctx,&pb.Employee{Name: "naga",Role: "develoveper"})
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println(updated_employee)

	de,err := c.DeleteEmployee(ctx,&pb.Employee{Name: "naga"})
	if err!=nil{
		fmt.Println(de)
		panic(err.Error())
	}
}