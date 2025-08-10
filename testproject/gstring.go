package main

import (
	"fmt"
	"strings"
	"strconv"
)


func main() {
	string_func()
	strconv_func()
}



func string_func(){
	strings.HasPrefix("test", "te")
	strings.HasSuffix("test", "st")
	strings.Contains("test", "es")
	strings.LastIndex("test","e")
	strings.IndexRune("test", 'e')
	strings.Replace("test", "e", "a", 1)
	strings.Repeat("tesssssst",3)
	strings.ToLower("TEST")
	strings.ToUpper("test")
	strings.TrimSpace(" test ")
	strings.TrimLeft(" test ", "t")
	strings.TrimRight(" test ", "t")
	strings.TrimPrefix(" test ", "t")
	strings.TrimSuffix(" test ", "t")
	strings.Trim(" test ", "t")
	strings.Fields("a b c d")
	strings.Split("testtesttest", "test")
	strings.Join([]string{"abc"}, "ddddfff")

	fmt.Println(strings.IndexRune("test", 'e'))
	fmt.Println(strings.Index("test", "es"))
	fmt.Println(strings.LastIndex("test", "e"))
	fmt.Println(strings.IndexRune("test", 'e'))
	fmt.Println(strings.IndexRune("test", 'e'))
	fmt.Println(strings.Replace("test", "e", "a", 1))
	fmt.Println(strings.Repeat("tesssssst",3))
	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))
	fmt.Println(strings.TrimSpace(" test "))
	fmt.Println(strings.TrimLeft(" test ", "t"))
	fmt.Println(strings.TrimRight(" test ", "t"))
	fmt.Println(strings.TrimPrefix(" test ", "t"))
	fmt.Println(strings.TrimSuffix(" test ", "t"))
	fmt.Println(strings.Trim(" test ", "t"))
	fmt.Println(strings.Fields("a b c d"))
	fmt.Println(strings.Split("testtesttest", "test"))
	fmt.Println(strings.Join([]string{"abc"}, "ddddfff"))



}

func strconv_func() {
	//strconv.IntSize()
	

	fmt.Println(strconv.IntSize)
	fmt.Println(strconv.Itoa(32))
	fmt.Println(strconv.Atoi("33"))


}