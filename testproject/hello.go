package main
import (
	"fmt"
	"os"
	"runtime"
)
var a = "G"

func init(){
	var (
		HOME = os.Getenv("HOME")
		USER = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)
	fmt.Println(HOME, USER, GOROOT)
}

func main(){
	//goos()
	
	fmt.Println("hello world")
}


func goos(){
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	fmt.Printf("path is %v\n", path)
}
