// mention the package - main is the main package which acts like entry point
// if there are any unused variables, the file wont run - throws compile time error

// if package level variables are declared using lowercase first letter then the variable is package level scoped
// if package level variables are declared using uppercase first letter then the variable is exposed to other packages
// if a variable is declared in block , it is block scoped but it will still be exposed if it starts with capital letter 

package main

//import "fmt"
var name string= "teja" //However cant use the := syntax
const (
	a=1+iota // default iota value starts at 0
	b // b is also assigned same value as 1+iota, however now iota is 1 as it acts like a counter
	c
)
var (
	name1 string="teja"
	age int= 21
)

func main() {
	//variable declaration 
	// var i int=42
	// j := 27 // automatically infers the type
	// k := 99.
	// name = "robert"
	// fmt.Printf("j %v, %T\n",j,j)
	// fmt.Printf("k %v, %T\n",k,k)
	// fmt.Printf("i %v, %T\n",i,i)
	// fmt.Printf("name %v, %T\n",name,name) // This is given precedence over package level variables - called as SHADOWING
	// fmt.Printf("name1 %v, %T\n",name1,name1)
	// fmt.Printf("age %v, %T\n",age,age)
	// fmt.Println("a b c\n",a,b,c)
	// k = float64(j) //type casting 
	// fmt.Printf("after type casting j to float64, k is %v, %T\n",k,k)
	// Arrays()
	// Pointers()
	// Functions("hey",123)
	// func(){
	// 	msg := "this is an anonoymous function. will execute automatically"
	// 	fmt.Println(msg)
	// }()
	// f := func(){
	// 	msg := "this function is assigned to a variable"
	// 	fmt.Println(msg)
	// 	fmt.Println(string([]rune("Z")[0]+32))
	// }
	// f()
	Caller()
	// Routines()
	// Mutex()
}