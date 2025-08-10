package main
import (
	f "fmt"
)

func main() {
	pointer_func()
	pointer_func2()
	pointer_func3()
}


func pointer_func() {
	var i1 = 5
	f.Println("i1 =%d is location in memeory is :%p\n", i1, &i1)

	var i2 = &i1
	f.Println("i2 =%d is location in memeory is :%p\n", i2, &i2)

	var i3 = *i2
	f.Println("i3 =%d is location in memeory is :%p\n", i3, &i3)

}

func pointer_func2() {
	var i1 = 10
	var intP *int
	intP = &i1

	f.Println("i1 =%d is location in memeory is :%p\n", i1, &i1)
	f.Println("*intP =%d intP = %d is location in memeory is :%p\n", *intP, intP, &intP)

}


func pointer_func3() {
	s := "good bye"
	var p * string = &s
	*p = "ciao"
	f.Println("here is pointerp: %p\n",p)
	f.Println("here is string s: %s\n",*p)
	f.Println("here is string s: %s\n",s)
}
