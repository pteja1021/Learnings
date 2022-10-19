package main
func (u User) TableName() string { // the TableName method can be used to give name to table. since User struct is passed as parameter, 
									// it will map to that table
	return "stakeholders"
}