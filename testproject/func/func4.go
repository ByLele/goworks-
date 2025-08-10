package main

import (
	"fmt"
	"log"
	"runtime"	
)

func main() {
	fmt.Println("hello world")
	
	// Define the where function inside main
	where := func() {
		_, file, line, ok := runtime.Caller(1)
		log.Printf("%s:%d %v", file, line, ok)
	}
	
	// Call the where function
	where()
	where()
}