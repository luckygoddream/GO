package main

import "fmt"

//go范围  范围可对各种数据结构中的元素进行迭代
func main() {
	//var num  = [3]string{"1","2","3"}
	nums := []int{1, 2, 3}
	//fmt.Println("num:",num)
	fmt.Println("nums:", nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
