package main

import "fmt"

var y string

type Rope string //类型别名

func main() {
	y = "G"
	print(y)
	fx()
	var str Rope = "Hello World"
	fmt.Println(str)
}

func fx() {
	y := "O"
	print(y)
	f3()
}

func f3() {
	print(y)
}
