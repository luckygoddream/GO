package main

import "fmt"

func main() {
	min(1, 3, 4, 7)
}
func min(s ...int) {
	for x, v := range s {
		fmt.Println(v)
		fmt.Println(x) //下标位置？

	}
}
