package main

import "fmt"

//go指针实例2
type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Person{"Bob", 20})

	fmt.Println(Person{name: "Alice", age: 30})

	fmt.Println(Person{name: "Jhon"})

	fmt.Println(&Person{name: "zhansan", age: 40})

	s := Person{name: "Ann", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
}
