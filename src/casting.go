package main

import (
	"fmt"
	"math"
)

func main() {
	var n int16 = 34
	var m int32
	m = int32(n)
	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
	fmt.Println(math.MinInt32)
	var x float64 = 3.11111111111
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		fmt.Println(int(whole))
	}
	//panic(fmt.Sprintf("%g is out of the int32 range", x))
}
