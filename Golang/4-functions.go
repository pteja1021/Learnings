package main

import (
	"fmt"
)
func Functions(parameter1 string,parameter2 int){
	fmt.Println("-----Functions-----")
	fmt.Println("Passed args are",parameter1,parameter2)
	fmt.Println(add(1,2))
	fmt.Println(varargs_function_sum(1,2,3,4,5))
	d,err := divide(5.0,0.0)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(d)
}
func add(num1, num2 int) int{ //means num1 and num2 both are ints . The end int is the return type
	return num1+num2
}
func varargs_function_sum(values ...int) int{ // the vararg param must be placed in the end
	fmt.Println("values will work like an array here. Values is ",values)
	sum := 0
	for _,v := range values{
		fmt.Println(v)
		sum += v
	}
	return sum 
}
func divide(a,b float32) (float32,error){
	if b==0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a/b, nil
}