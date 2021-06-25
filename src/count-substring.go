package main

import (
	"fmt"
	"strings"
)

//Count 用于计算字符串 str 在字符串 s 中出现的非重叠次数：strings.Count(s, str string) int
func main() {
	str := "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"
	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}
