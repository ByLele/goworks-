package main

import (
	"fmt"
	"time"
)


func main() {
	date_time_func()
}


func date_time_func() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(time.Now().Location())
	fmt.Println(time.Now().Zone())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().After(time.Now().Add(-time.Hour)))
	fmt.Println(time.Now().Before(time.Now().Add(time.Hour)))
	fmt.Println(time.Now().Equal(time.Now()))
	fmt.Println(time.Now().IsZero())
	fmt.Println(time.Unix(time.Now().Unix(), 0))
	fmt.Println(time.Unix(0, time.Now().UnixNano()))
	fmt.Println(time.NewTicker(time.Second))
}