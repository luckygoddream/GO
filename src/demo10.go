package main

import "fmt"

//go的闭包（匿名函数实例）
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	nowInt := intSeq()
	fmt.Println(nowInt())
	fmt.Println(nowInt())
	fmt.Println(nowInt())
	newInt := intSeq()
	fmt.Println(newInt())

}
