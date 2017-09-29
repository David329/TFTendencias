package main

import (
	"sync"
	"log"
)

// Parallelize parallelizes the function calls
func Parallelize(functions ...func()) {
    var waitGroup sync.WaitGroup
    waitGroup.Add(len(functions))

    defer waitGroup.Wait()

    for _, function := range functions {
        go func(copy func()) {
            defer waitGroup.Done()
            copy()
        }(function)
    }
}

func main(){
	func1:= func() {
		for i:=0;i<10;i++{
			log.Printf("A: %d",i)
		}
	}
	
	func2:= func() {
		for i:=0;i<10;i++{
			log.Printf("B: %d",i)
		}
	}
	
	func3:= func() {
		for i:=0;i<10;i++{
			log.Printf("C: %d",i)
		}
	}
	
	Parallelize(func1, func2, func3)
}