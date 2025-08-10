package main

import (
	"fmt"
)


func main() {
	range_byte()

}


func range_byte() {
    s := "\u00ff\u754c"
    for i, c := range s {
        fmt.Printf("%d:%c ", i, c)
    }

}


func Compare(a,b[]byte) int {
    for i:= 0;i < len(a)&&i < len(b);i++{
        switch {
        case a[i]>b[i]:
            return 1
        case a[i]<b[i]:
            return -1
        }
    }
    switch {
    case len(a) <len(b):
        return -1
    case len(a)>len(b):
        return 1
    }
    return  0
}