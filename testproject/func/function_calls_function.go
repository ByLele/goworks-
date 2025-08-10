package main
import (
	"fmt"
	"math/rand"
    "time"
)

var a string

func main() {
    RandomNum()
}

func uint8fromint () {
    
	a, err := Uint8FromInt(256)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
func f1() {
   a := "O"
   print(a)// a为 f1函数中的app
   f2()// a为 main函数中的a
}

func f2() {
   print(a)
}

func n(){print(a)}

func m(){
	a :='O'
	print(a)
}

func local_scope1(){
	n()
	m()
	n()
}


func Uint8FromInt(n int) (uint8, error) {
    if 0 <= n && n <= math.MaxUint8 { // conversion is safe
        return uint8(n), nil
    }
    return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func IntFromFloat64(x float64) int {
    if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
        whole, fraction := math.Modf(x)
        if fraction >= 0.5 {
            whole++
        }
        return int(whole)
    }
    panic(fmt.Sprintf("%g is out of the int32 range", x))
}

func RandomNum() {
    for (i:=0;i<10;i++){
        a := rand.Int()
        fmt.Printf("%d/ ",a)
    }

    for (i:=0;i<5;i++){
        r := rand.Intn(8)
        fmt.Printf("%d / ",r)
    }
    fmt.Println()
    timens := int64(time.Now().Nanosecond())
    rand.Seed(timens)
    for(i:=0;i<10;i++){
        fmt.Printf("%2.2f / ",100*rand.Float32())
    }
}

