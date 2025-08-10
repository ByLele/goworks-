package main

import "fmt"

func main() {
	for_arrays()
	for_string()
	pointer_array()
	array_value()
	multidim_array()
	array := [3]float64{7.0,8.5,9.1}
	x:= array_sum(&array)

	fmt.Printf("The sum is: %f\n",x)
	array_slices()
	string_slice()
	slice_sum2()
	make_slice()

}




func for_arrays() {
	var arr1[5] int
	for i:=0;i<len(arr1);i++{
		arr1[i] = i*2
	}
	fmt.Println(arr1)


	for i:=0;i<len(arr1);i++{
		fmt.Printf("array at index %d is %d\n",i,arr1[i])
	}


	for i,v:=range arr1{
		fmt.Printf("array at index %d is %d\n",i,v)
	}

}

func for_string() {
	a:= [...]string{"a","b","c","d","e"}
	for i:=0;i<len(a);i++{
		fmt.Printf("array at index %d is %s\n",i,a[i])
	}

	for i,v:=range a{
		fmt.Printf("array at index %d is %s\n",i,v)
	}
	
}
func f(a [3]int) {fmt.Println(a)}

func fp(a *[3]int) {fmt.Println(a)}

func pointer_array() {
	var ar [3]int
	f(ar)
	fp(&ar)
}

func for_array() {
	var arr1 [15]int
	for i:=0;i<len(arr1);i++{
		arr1[i] = i*2
	}

	for i:=0;i<len(arr1);i++{
		fmt.Printf("array at index %d is %d\n",i,arr1[i])
	}
	
}

func fp3(a *[3]int) { fmt.Println(a)}

func array_value() {
	var arrKeyValue = [5]string{3:"Chris",4:"Ron"}
	for i:=0;i<len(arrKeyValue);i++{
		fmt.Printf("Person at %d is %s\n",i,arrKeyValue[i])
	}

	arr3 := [...]int{1,2,3}
	for i:=0;i<len(arr3);i++{
		fmt.Printf("Person at %d is %s\n",i,arr3[i])
	}


	for i:=0;i<3;i++ {
		fp3(&[3]int{i,i*2,i*i*3})
	}
}

const WIDTH int= 1920
const HEIGH int= 1080
type pixel int

var screen [WIDTH][HEIGH]pixel
func multidim_array() {

	for y:=0;y<HEIGH;y++ {
		for x:=0;x<WIDTH;x++ {
			screen[x][y] = 0
		}
	}

}



func array_sum(a *[3]float64) (sum float64) {
	for _,v:= range *a {
		sum += v
	}	
	return
}

func array_slices() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]
	
	for i:=0;i<len(slice1);i++ {
		fmt.Printf("Slice at %d is %d\n",i,slice1[i])
	}

	for i:=0;i<len(arr1);i++ {
		fmt.Printf("Array at %d is %d\n",i,arr1[i])
	}
	fmt.Printf("The length of arr1 is %d\n",len(arr1))
	fmt.Printf("The length of slice1 is %d\n",len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n",cap(slice1))

	slice1 = slice1[1:4]
	for i:=0;i<len(slice1);i++ {
		fmt.Printf("slice1 at %d is %d\n",i,slice1[i])
	}

	fmt.Printf("The length of slice1 is %d\n",len(slice1))
	fmt.Printf("The cap of slice1zzzzz is %d\n",cap(slice1))
}



func string_slice() {
	b := []byte{'g','o','l','a','n','g'}
	fmt.Println(b[1:4])

	fmt.Println(b[:2])

	fmt.Println(b[2:])

	fmt.Println(b[:])



}

func slice_sum(a []int) int {
	s:=0
	for i:=0;i<len(a);i++ {
		s += a[i]
	}
	return s
}
func slice_sum2() int {
	var arr = [5]int{1,2,3,4,5}
	fmt.Println(slice_sum(arr[:]))
	return slice_sum(arr[:])
}

func make_slice() {
	var slice1 []int = make([]int,10)
	for i:=0;i<len(slice1);i++{
		slice1[i] = i*5
	}
	for i:=0;i<len(slice1);i++{
		fmt.Printf("Slice at %d is %d\n",i,slice1[i])
	}

	slice2 := make([]int,10,20)
	for i:=0;i<len(slice2);i++{
		slice2[i] = i*5
	}
	
}