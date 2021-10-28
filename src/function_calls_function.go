package main

var xa string

func main() {
	xa = "G"
	print(xa)
	fa()
}
func fa() {
	xa := "0"
	print(xa)
	fb()
}

func fb() {
	print(xa)
}
