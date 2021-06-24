package main

var y string

func main() {
	y = "G"
	print(y)
	fx()
}

func fx() {
	y := "O"
	print(y)
	f3()
}

func f3() {
	print(y)
}
