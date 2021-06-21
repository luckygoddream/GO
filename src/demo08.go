package main

import "fmt"

//Go函数多个返回值实例
func vals() (int, int) {
	return 3, 7
}
func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
