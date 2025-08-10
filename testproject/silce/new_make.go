package main

import (
	"fmt"
	"bytes"
)


func main() {
	make1()
	copy_append_slice()

}

func make1() {
	s := make([]int,5)
	fmt.Println(len(s),cap(s))
	s = s[2:4]
	fmt.Println(len(s),cap(s))
	s1 := []byte{'g','o','l','a','n','g'}
	fmt.Println(s1)
	s2 := s1[2:]
	fmt.Println(s2)
	s2[1] = 't'
	fmt.Println(s1)
	fmt.Println(s2)

}

func for_range() {
	var slice1 []int = make([]int,10)
	slice1[0] = 4
	slice1[1] = 5
	slice1[2] = 6

	for i,v := range slice1 {
		fmt.Printf("Slice at %d is %d\n",i,v)
	}


	seasons := []string{"Spring","Summer","Autumn","Winter"}
	for i,v := range seasons {
		fmt.Printf("Slice at %d is %s\n",i,v)
	}

	for _,v := range seasons {
		fmt.Printf("Slice is %s\n",v)
	}

		
}


func copy_append_slice() {
	sl_from := []int{1,2,3}
	sl_to := make([]int,10)
	n := copy(sl_to,sl_from)
	fmt.Println(sl_to)
	fmt.Printf("copied %d elements \n",n)
	sl3 := []int{1,2,3}
	sl3 = append(sl3,4,5,6)
	fmt.Println((sl3))

}


func AppendByte(slice []byte,data...byte) []byte{
	m := len(slice)
	n := m+len(data)
	if n > cap(slice) {
		newSlice := make([]byte,(n+1)*2)
		copy(newSlice,slice)
		slice = newSlice

	}
	slice = slice[0:n]
	copy(slice[m:n],data)
	return slice

}


func magnify_slice(slice []int,factor int) []int {

	if factor <=0 {
		return []int {}
	}
		
	n :=  len(slice)*factor
	newSlice := make([]int,n)
	copy(newSlice,slice)
	slice = newSlice
	return slice
}
/*
	filter_slice.go

使用高阶函数对一个集合进行过滤：s 是前 10 个整数的一个切片。建立一个函数 Filter，它接受 s 作为第一参数，fn func (int) bool 作为第二参数，并返回满足函数 fn 的 s 元素的切片（使其为真）。用 fn 测试整数是否为偶数。
*/
func filter_slice(s []int,fn func(int)bool) []int{
	result := make([]int,0)
	for _,v := range s{
		if fn(v){
			result = append(result,v)
		}
	}
	return result

}


func isEven(n int) bool{
	return n%2 == 0
}


/*

insert_slice.go

编写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。

*/

func insert_slice(newSlice[]int,oldSlice[]int,index int) []int {

	if len(newSlice) <= len(oldSlice) {
		return []int{}
	}
	if index <0 || index > len(newSlice) {
		return []int {}
	}

	targetSlice := make([]int,0,len(newSlice)+len(oldSlice))
	targetSlice = append(targetSlice,newSlice[:index]...)
	targetSlice = append(targetSlice,oldSlice...)
	targetSlice = append(targetSlice,newSlice[index:]...)
	return targetSlice

}

/*
emove_slice.go

编写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除。
*/
func remove_slice(slice []int,start,end int) []int {

	targetSlice := make([]int,len(slice)-(end -start))
	targetSlice = append(targetSlice,slice[:start]...)
	targetSlice = append(targetSlice,slice[end:]...)
	return targetSlice
}