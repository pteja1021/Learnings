// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/employee.proto

package employeeDB

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Role         string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	ManagerId    int64  `protobuf:"varint,4,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
	DepartmentId int64  `protobuf:"varint,5,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Employee) GetManagerId() int64 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

func (x *Employee) GetDepartmentId() int64 {
	if x != nil {
		return x.DepartmentId
	}
	return 0
}

type NewEmployee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Role         string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	ManagerId    int64  `protobuf:"varint,3,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
	DepartmentId int64  `protobuf:"varint,4,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
}

func (x *NewEmployee) Reset() {
	*x = NewEmployee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewEmployee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewEmployee) ProtoMessage() {}

func (x *NewEmployee) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewEmployee.ProtoReflect.Descriptor instead.
func (*NewEmployee) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{1}
}

func (x *NewEmployee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewEmployee) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *NewEmployee) GetManagerId() int64 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

func (x *NewEmployee) GetDepartmentId() int64 {
	if x != nil {
		return x.DepartmentId
	}
	return 0
}

type EmptyEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyEmployeeRequest) Reset() {
	*x = EmptyEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyEmployeeRequest) ProtoMessage() {}

func (x *EmptyEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyEmployeeRequest.ProtoReflect.Descriptor instead.
func (*EmptyEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{2}
}

type EmptyEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyEmployeeResponse) Reset() {
	*x = EmptyEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyEmployeeResponse) ProtoMessage() {}

func (x *EmptyEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyEmployeeResponse.ProtoReflect.Descriptor instead.
func (*EmptyEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{3}
}

type AllEmployees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employees []*Employee `protobuf:"bytes,1,rep,name=employees,proto3" json:"employees,omitempty"`
}

func (x *AllEmployees) Reset() {
	*x = AllEmployees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_employee_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllEmployees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllEmployees) ProtoMessage() {}

func (x *AllEmployees) ProtoReflect() protoreflect.Message {
	mi := &file_proto_employee_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllEmployees.ProtoReflect.Descriptor instead.
func (*AllEmployees) Descriptor() ([]byte, []int) {
	return file_proto_employee_proto_rawDescGZIP(), []int{4}
}

func (x *AllEmployees) GetEmployees() []*Employee {
	if x != nil {
		return x.Employees
	}
	return nil
}

var File_proto_employee_proto protoreflect.FileDescriptor

var file_proto_employee_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x44, 0x42, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x0b, 0x4e,
	0x65, 0x77, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17,
	0x0a, 0x15, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x44, 0x42, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x52, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x32, 0xb7, 0x02, 0x0a, 0x14,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x43, 0x72, 0x75, 0x64, 0x12, 0x3e, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x12, 0x17, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x44, 0x42,
	0x2e, 0x4e, 0x65, 0x77, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x1a, 0x14, 0x2e, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x44, 0x42, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x44,
	0x42, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x44, 0x42, 0x2e, 0x41, 0x6c, 0x6c, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x4f, 0x66, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x14, 0x2e, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x44, 0x42, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x1a, 0x14, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x44, 0x42, 0x2e, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x14, 0x2e, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x44, 0x42, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x1a, 0x21, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x44, 0x42, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x44, 0x42, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_employee_proto_rawDescOnce sync.Once
	file_proto_employee_proto_rawDescData = file_proto_employee_proto_rawDesc
)

func file_proto_employee_proto_rawDescGZIP() []byte {
	file_proto_employee_proto_rawDescOnce.Do(func() {
		file_proto_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_employee_proto_rawDescData)
	})
	return file_proto_employee_proto_rawDescData
}

var file_proto_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_employee_proto_goTypes = []interface{}{
	(*Employee)(nil),              // 0: employeeDB.Employee
	(*NewEmployee)(nil),           // 1: employeeDB.NewEmployee
	(*EmptyEmployeeRequest)(nil),  // 2: employeeDB.EmptyEmployeeRequest
	(*EmptyEmployeeResponse)(nil), // 3: employeeDB.EmptyEmployeeResponse
	(*AllEmployees)(nil),          // 4: employeeDB.AllEmployees
}
var file_proto_employee_proto_depIdxs = []int32{
	0, // 0: employeeDB.AllEmployees.employees:type_name -> employeeDB.Employee
	1, // 1: employeeDB.EmployeeDatabaseCrud.AddEmployee:input_type -> employeeDB.NewEmployee
	2, // 2: employeeDB.EmployeeDatabaseCrud.GetEmployees:input_type -> employeeDB.EmptyEmployeeRequest
	0, // 3: employeeDB.EmployeeDatabaseCrud.UpdateRoleOfEmployee:input_type -> employeeDB.Employee
	0, // 4: employeeDB.EmployeeDatabaseCrud.DeleteEmployee:input_type -> employeeDB.Employee
	0, // 5: employeeDB.EmployeeDatabaseCrud.AddEmployee:output_type -> employeeDB.Employee
	4, // 6: employeeDB.EmployeeDatabaseCrud.GetEmployees:output_type -> employeeDB.AllEmployees
	0, // 7: employeeDB.EmployeeDatabaseCrud.UpdateRoleOfEmployee:output_type -> employeeDB.Employee
	3, // 8: employeeDB.EmployeeDatabaseCrud.DeleteEmployee:output_type -> employeeDB.EmptyEmployeeResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_employee_proto_init() }
func file_proto_employee_proto_init() {
	if File_proto_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewEmployee); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_employee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyEmployeeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_employee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyEmployeeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_employee_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllEmployees); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_employee_proto_goTypes,
		DependencyIndexes: file_proto_employee_proto_depIdxs,
		MessageInfos:      file_proto_employee_proto_msgTypes,
	}.Build()
	File_proto_employee_proto = out.File
	file_proto_employee_proto_rawDesc = nil
	file_proto_employee_proto_goTypes = nil
	file_proto_employee_proto_depIdxs = nil
}
