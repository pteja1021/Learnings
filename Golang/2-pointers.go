package main
import "fmt"
func Pointers(){
	var a int = 42 
	var b *int = &a
	fmt.Println("a,b",a,b) // b gives the address of a
	fmt.Println("a,*b",a,*b) //*b dereferences it and gives the value at b's address
	a=27
	fmt.Println("after changing a to 27, the values of a, *b is",a,*b)

	array := [...]int{1,2,3}
	array_b := &array[0]
	array_c := &array[1]
	fmt.Println("array,array_b,array_c",array,array_b,array_c) // addresses of array_b and array_c are 4 bytes apart since int is 4bytes 
}

// if a pointer is not initialised, default value of nil is given 