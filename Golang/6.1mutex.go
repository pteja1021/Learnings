package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg1 = sync.WaitGroup{}
var counter =0 
var m = sync.RWMutex{} // Read Write Mutex

func Mutex(){
	runtime.GOMAXPROCS(100) // sets the max number of CPUS that can be executing
	for i:=0; i<10; i++ {
		wg1.Add(2)
		m.RLock()
		go sayHello()
		m.Lock() // placing the locks here will prevent the race condition
		go increment()
	}
	wg1.Wait()
}

func sayHello() {
	// m.RLock() // places a read lock
	// however placing the locks inside the threads will stop the counter from being simultaneouly changed but wont synchronise the threads
	// the way we want to. 
	fmt.Println("hello ",counter)
	m.RUnlock() //unlock the read lock
	wg1.Done() // thread completion
}

func increment(){
	// m.Lock() // places the write lock
	counter++
	m.Unlock() // unlock the write lock
	wg1.Done() // thread completion
}