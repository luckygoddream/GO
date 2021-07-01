package main

import (
	"fmt"
	"time"
)

func main() {
	star := time.Now()
	x := fib(5)
	fmt.Println(x)
	des(10)

	z, c := fib2(10)
	fmt.Println(z, c)

	m := muiltply(4)
	fmt.Println(m)
	end := time.Now()
	sub := end.Sub(star)
	fmt.Println(sub)
}
func fib(i int) int {
	if i <= 2 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func des(i int) int {
	fmt.Println(i)
	if i == 0 {
		return 0
	}
	return des(i - 1)
}

func fib2(i int) (x int, y int) {
	if i <= 2 {
		return i, 1
	}
	_, j := fib2(i - 1)
	_, k := fib2(i - 2)
	return i, j + k
}

func muiltply(i int) int {
	if i == 0 {
		return 1
	}
	return i * muiltply(i-1)
}
