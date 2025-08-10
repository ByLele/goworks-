package main

import "fmt"
type Stack struct {
	ix int
	data [10]int
}

func main() {
	fmt.Println("Hello, World!")
	greeting()
	fmt.Println(simple_func(1,2,3))
	numx2,numx3 = getX2AndX3(num)
	PrintValue()
	numx2,numx3 = getX2And3_2(num)
	PrintValue()

	PrintThreeValues()
	PrintMinMax()
	PrintSideEffect()
	varnumpar()
	defer1()	
}

func greeting() {
	fmt.Println("in greeting")
	//greeting() //无限递归
	greeting_call()
	fmt.Println("in greeting")
}

func greeting_call() {
	fmt.Println("In Greeting: Hi!!!!!!")
}

func (st *Stack) Pop() int {
	v := 0
	for ix := len(st.data)-1; ix >= 0; ix-- {
		if v == st.data[ix] && v != 0 {
			st.data[ix] = 0
			return v
		}
	}
	return 0
}
var num int = 10
var numx2,numx3 int
func simple_func(a,b,c int) (int) {
	return a*b*c
}

func multiple_return() {

}

func PrintValue() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n",num,numx2,numx3)
}

func getX2AndX3(input int) (int,int) {
	return 2*input,3*input
}

func getX2And3_2(input int) (x2 int,x3 int) {
	x2 = 2*input
	x3 = 3*input
	return
}

func PrintThreeValues() {
	var i1,i2 int
	var i3 float32

	i1,i2,i3 = ThreeValues()
	fmt.Printf("Three values: %d,%d,%f\n",i1,i2,i3)

	a, b, c := ThreeValues2()
	fmt.Printf("Three values: %d,%d,%d\n", a, b, c)
}	

func ThreeValues() (int,int,float32) {
	return 5, 6, 7.5
}

func ThreeValues2() (a int,b int,c int) {
	return 5,6,7
}
func PrintMinMax() {
	min,max := MinMax(98,90)
	fmt.Printf("minmum is %d, maximum is %d\n",min,max)
}

func MinMax(a,b int) (min int,max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}	

func PrintSideEffect() {
	n := 0
	reply := &n
	side_effect(1,2,reply)
	fmt.Println("n = ",n) //n 地址replay 存储 3
	fmt.Println("reply = ",reply)
}

func side_effect(a,b int,replay* int) {
	*replay = a + b
}
func varnumpar() {
	nums := []int{1,2,3,4,5}
	fmt.Println("min = ",min(nums...))
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _,v := range s {
		if v < min {
			min = v
		}
	}	
	return min	
}
func defer1() {
	fmt.Println("defer1")
	defer defer2()
	fmt.Println("defer1 end")
}
func defer2() {
	fmt.Println("defer2")
}
