package main

import "fmt"

func f111(a [3]int) { fmt.Println(a) }
func fp(a *[3]int)  { fmt.Println(a) }

func main() {
	var ar [3]int
	f111(ar) // passes a copy of ar
	fp(&ar)  // passes a pointer to ar
}
