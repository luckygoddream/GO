package main

import "fmt"

//defer实例 与java中finally相似
func main() {
	f11()
	aaa()
	bbb()
}
func f11() {
	fmt.Println("Hello f1")
	defer fy()
	fmt.Println("world f1")
}

func fy() {
	fmt.Println("hello f2")
}

func aaa() {
	i := 0
	defer fmt.Println(i) //输出0
	i++
	return
}
func bbb() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) //倒序输出类似栈，先进后出，后进先出
	}
}
