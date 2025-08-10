package main

import (
	"fmt"
	"strings"
)
// 闭包 匿名函数

func main() {
	p2 := add2()
	fmt.Println("Call add2 for 3 gives:%v\n",p2(3))

	TwoAdder := Adder(2)
	fmt.Println("Call Adder(2) for 3 gives:%v\n",TwoAdder(3))
	ThreeAdder := Adder(3)
	fmt.Println("Call Adder(3) for 4 gives:%v\n",ThreeAdder(4))
	fmt.Println("Call Adder(3) for 4 gives:%v\n",ThreeAdder(4))
	func_close()

	fib := fibonacci()
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", fib())
	}
	fmt.Println()

	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	fmt.Println(addBmp("file"))
	fmt.Println(addJpeg("file"))

}

// func add2() (func (b int) int)
// func Adder(a int) (func (b int) int)

func add2() (func (b int) int) {
	return func(b int) int {
		return b+2
	}
}

func Adder(a int) (func (b int) int) {
	return func(b int) int {
		return a + b
	}
}	


func func_close() {
	var f = Adder2()
	fmt.Print(f(1),"\n")
	fmt.Print(f(10),"\n")
	fmt.Print(f(100),"\n")

}

func Adder2() func(int) int {
	var x int
	return func(delta int) int {
		x+= delta
		return x
	}
}


func fibonacci() func () int {
	a,b := 0,1
	return func() int {
		result := a
		a,b= b,a+b
		return result
	}
}


func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name,suffix) {
			return name + suffix
		}
		return name
	}
}