package main

import (
	"context"
	"log"
	"net"

	model "example.com/employeeDB/models"
	pb "example.com/employeeDB/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

)

const (
	port = ":50051"
)

type employeeDBServer struct{
	pb.UnimplementedEmployeeDatabaseCrudServer
	db *gorm.DB // db connection will be part of this struct and can be accesible
}

func (s *employeeDBServer) AddEmployee(ctx context.Context, in *pb.NewEmployee) (*pb.Employee, error) {
	log.Printf("Received %v %v %v",in.GetName(),in.GetDepartmentId(),in.GetRole())
	newEmployee := model.Employee{
		Name: in.GetName(),
		Role: in.GetRole(),
		DepartmentID: uint(in.GetDepartmentId()),
	}
	s.db.Save(&newEmployee)
	return &pb.Employee{Name: in.GetName(), Role: in.GetRole(), DepartmentId: in.GetDepartmentId(),Id: int64(newEmployee.ID)},nil
}

func (s* employeeDBServer) GetEmployees(ctx context.Context, in *pb.EmptyEmployeeRequest) (*pb.AllEmployees, error) {
	log.Println("Get All employees request received")
	allEmployees := []model.Employee{}
	allEmployeesResult := []*pb.Employee{}
	s.db.Find(&allEmployees)
	for _,employee := range allEmployees{
		allEmployeesResult = append(allEmployeesResult, &pb.Employee{Name: employee.Name, Role: employee.Role, DepartmentId: int64(employee.DepartmentID),Id: int64(employee.ID)})
	}
	return &pb.AllEmployees{Employees: allEmployeesResult},nil
} 

func (s* employeeDBServer) UpdateRoleOfEmployee(ctx context.Context,in *pb.Employee) (*pb.Employee, error){
	log.Println("Update employee request received")
	s.db.Model(&model.Employee{}).Where("name=?",in.Name).Update("role",in.Role)
	return &pb.Employee{Name: in.GetName(),Role: in.GetRole()},nil
}

func (s* employeeDBServer) DeleteEmployee(ctx context.Context, in *pb.Employee) (*pb.EmptyEmployeeResponse,error){
	log.Println("Delete employee request received")
	s.db.Where(&model.Employee{Name: in.GetName()}).Delete(&model.Employee{})
	return &pb.EmptyEmployeeResponse{},nil
}

func main(){
	model.InitiateDB()
	//Open port
	lis, err := net.Listen("tcp",port)
	if err!=nil{
		log.Fatal(err.Error())
	}

	//creating db connection
	conn,err := gorm.Open("postgres","user=postgres password=bcdev123 dbname=gorm sslmode=disable")
	if err!=nil{
		panic(err.Error())
	}
	defer conn.Close()

	// registering our server
	s:= grpc.NewServer()
	
	pb.RegisterEmployeeDatabaseCrudServer(s,&employeeDBServer{
		db: conn,
	})

	log.Printf("Listening to server on %v",lis.Addr())
	
	// Serving
	if err := s.Serve(lis); err!=nil {
		log.Fatal("Unable to serve",err.Error())
	}
}
