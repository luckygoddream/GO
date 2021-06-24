# GO学习小结
## 1.int 和 int64 是相同的类型吗？
    package main
    func main() {
    var a int
    var b int64
    a = 15
    // b = a + a 	// 编译错误： cannot use a + a (type int) as type int64 in assignment
    b = a + 5 	// 5
    }
    在执行 b = a + a 的时候会报编译错误，也就是说 a 和 b 是不同的类型。
