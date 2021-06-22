package main

import "fmt"

//go的指针实例1
func zeroVal(iVal int) {
	iVal = 0
}
func zeroPre(iptr *int) {
	*iptr = 0
}
func main() {
	var i int = 1
	fmt.Println("initial", i)
	zeroVal(i)
	fmt.Println("zeroval", i)
	zeroPre(&i)
	fmt.Println("zeroptr", i)
	fmt.Println("pointer", &i)

}
