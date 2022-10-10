package main
import "fmt"
func Arrays(){
	a1 := []int{1,2,3,4,5,6,7}
	b := a1[:]
	c := a1[2:]
	d := a1[:3]
	e := a1[3:6]
	f := e //passed by reference ,this will also affect a1,b,c,d along with e since all are passed by references 
	f[1] =4
	fmt.Println("a1",a1)
	fmt.Println("b",b)
	fmt.Println("c",c)
	fmt.Println("d",d)
	fmt.Println("e",e)
	fmt.Println("f",f)
}