package main

//go的常量
import (
	"fmt"
	"math"
)

func main() {
	const name = "张三"
	fmt.Println(name)
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
