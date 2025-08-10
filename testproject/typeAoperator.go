package main

import "fmt"
import "strings"

type TZ int
type MS string

func main() {
	char()
}

func types() {
	var a,b int = 3,4
	c := a +b
	fmt.Println(c)	
	var ms MS = "hello"
	fmt.Println(ms)
	var s string = "WORLD"
	fmt.Println(strings.ToLower(s))
	//fmt.Println(s.lower())
}

func char() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3) // UTF-8 code point
}



