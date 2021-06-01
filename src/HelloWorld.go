package main

import "fmt"

func main() {
	var a string = "123"
	var x, y int = 1, 2
	/** intVal := 1 相当于 var intVal int = 1 变量已声明未使用会报错 定义常量与全局变量不使用不会报错
	**/
	intVal := 1
	f := "DARK"
	//const 定义常量
	const ss string = "111"
	const age = 30
	const name1, name2 string = "jack", "rusi"
	const age1, age2 = 14, 23
	const ff, num = false, 333 //多重赋值，自动推断类型
	fmt.Println(age1, age2)
	fmt.Println(name1, name2)
	fmt.Println(age)
	fmt.Println(ss)
	fmt.Println(f)
	fmt.Println(intVal)
	fmt.Println(x, y)
	fmt.Println(a)
	fmt.Println("Hello Word")
}
