package main

import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}

func Routines(){
	//go sayHello() 
	msg := "Inside Routines, outside synonymous functions"
	wg.Add(1) // adds 1 saying that a routine. 1 is count of threads being added
	go func(){
		fmt.Println(msg) // all outer variables are accessible to this func even though both are executing in different call stack
		wg.Done() // tells the wait group that it has completed its execution and decrements the no. of threads the waitgroup is waiting on
	}() // try passing the msg as argument and see the result
	
	msg = "After routine invocation"
	wg.Wait() // waits for the routines to complete. Basically exits the application
}

func SayHello(){
	fmt.Println("hello")
}