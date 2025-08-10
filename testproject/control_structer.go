package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	if_func()
	Abs(10)
	isGreater(10,20)
	num_big()
	string_conversion2()
	switch1()
	switch2()

	//for1()
	//for_string()
	for_loop()
	fizzbuzz()
	rectangle_stars()
}


func if_func() {
	bool1 := true
	if bool1 {
		fmt.Println("the value is true")
	} else {
		fmt.Println("the value is false")
	}
}

func init() {
	var prompt string = "Enter a digit, e.g. 3" + "or %s to quit ."
	if runtime.GOOS == "windows" {
		fmt.Println("platform is windows")
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	}else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}


}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isGreater (x,y int) bool {
	if x > y {
		return true
	}
	return false
}


func num_big() {
	var first int = 0
	var cond int
	if first <= 0 {
		fmt.Printf("first is less then or equal to 0\n")
	} else if first > 0 && first <5 {
		fmt.Printf("first is between 0 and 5\n")
	}else {
		fmt.Printf("first is 5 or greater")
	}

	if cond =5 ; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	}else {
		fmt.Printf("cond is less then or equal to 10\n")
	}


}

func string_conversion2() {
	var orig string = "ABC"
	var newS string
	fmt.Printf("the size of int is: %d \n",strconv.IntSize)
	an,err:= strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n",orig)
		return
	}
	fmt.Printf("The integer is %d\n",an)
	an = an +5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n",newS)

}

// switch 
func switch1() {
	var num1 int = 100
	switch num1 {
	case 98,99:
		fmt.Println("it`s equal to 98")
	case 100:
		fmt.Println("its equal to 100")
	default:
		fmt.Println("it`s not equal to 98 or 100")
	}
}

func switch2() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:		fmt.Println("was <= 5"); fallthrough;
	case 6:
		fmt.Println("was <= 6")
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

//for
func for1(){
	for i:=10;i< 100;i++ {
		fmt.Printf("the counter is at %d\n",i)
	}
}


func for_string(){ 
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n",len(str))
	for ix:= 0;ix<len(str);ix++{
		fmt.Printf("Character on position %d is: %c \n",ix,str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n",len(str2))
	for ix:=0;ix<len(str2);ix++ {
		fmt.Printf("Character on position %d is: %c \n",ix,str2[ix])
	}

}

func for_loop() {
	for i:=0;i<25;i++ {
		for j:=0;j<i;j++{
			fmt.Printf("G")
		}
		fmt.Printf("\n")
	}
}


func bitwise_complement() {
	var a uint8 = 0xAB
	var b uint8 = 0x55
	fmt.Printf("The result of %08b & %08b is %08b\n",a,b,a&b)
	fmt.Printf("The result of %08b | %08b is %08b\n",a,b,a|b)
	fmt.Printf("The result of %08b ^ %08b is %08b\n",a,b,a^b)
	fmt.Printf("The result of %08b &^ %08b is %08b\n",a,b,a&^b)
	fmt.Printf("The result of %08b << 1 is %08b\n",a,a<<1)
	fmt.Printf("The result of %08b >> 1 is %08b\n",a,a>>1)
	
}

func fizzbuzz(){
	for i:=1;i<=100;i++ {
		if i %3 == 0 {
			fmt.Printf("Fizz")
		}
		if i %5 == 0 {
			fmt.Printf("Buzz")
		}
		if i %3 != 0 && i %5 != 0 {
			fmt.Printf("%d",i)
		}
		fmt.Printf("\n")		
	}
}

func rectangle_stars() {
	for i:=0;i<10;i++ {
		for j:=0;j<10;j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}


func for6() {
LABEL1:
	for i := 0 ;i <=5;i++ {
		for j := 0;j<3;j++ {
			if j == 2 {
				continue LABEL1
			}
			fmt.Printf("i is %d, j is %d\n",i,j)
		}
	}
}


func goto()  {
	i := 0
	HERE:
		fmt.Println(i)
		i++
		if i == 5 {
			return
		}
		goto HERE
}