package main

import "fmt"
import "io"
import "log"

func main() {
	fmt.Println("hello world")
	defer3()
	defer4()
	defer_tracing()
	func1("go")
	fibonacci_test()
	mut_recurs()	
	func_parameter()	
	func_literal()
	fmt.Println(f())
}



func defer3() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func defer4() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ",i)
	}
}

func trace(s string) {fmt.Println("entering : %s \n",s)}
func untrace(s string) { fmt.Println("leaving:%d\n",s)}

func defer_tracing() {
	trace("defer_tracing")
	defer untrace("defer_tracing")
	fmt.Println("in defer_tracing")
}

func func1(s string) (n int,err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v\n",s,n,err)
	}()
	return 7,io.EOF
}



func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci_test() {
	fib := fibonacci(10)
	fmt.Println(fib)
}


func mut_recurs() {
	fmt.Println(even(17))
	fmt.Println(odd(17))
}

func even(n int) bool {
	if n == 0 {
		return true
	}
	return odd(resign(n-1))
}

func odd(n int) bool {
	if n == 0 {
		return false
	}
	return even(resign(n-1))
}	

func resign(n int) int {
	if n< 0 {
		return -n
	}
	return n
}

func func_parameter() {
	callback(1,Add)
}

func Add(a,b int)  {
	fmt.Println("Add%d+ %d = %d",a,b,a+b)
	
}

func callback(y int,f func(int,int) )  {
	f(y,2)
}

func func_literal() {
	func() {
		for i :=0;i<4;i++ {
			g := func(i int) {
				fmt.Printf("%d",i)
			}
			g(i)
			fmt.Printf("- g is of type %T and has value %v\n",g,g)
		}
	}()
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}