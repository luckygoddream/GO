package main

import "fmt"

func main() {
	var s [5]int
	for i := 0; i < len(s); i++ {
		s[i] = i * 2
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("Array at index %d is %d\n", i, s[i])
	}

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

}
