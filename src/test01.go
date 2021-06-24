package main

var a = "G"

func main() {
	n()
	l()
	m()
	n()

}

func n() { print(a) }

func m() {
	a := "O"
	print(a)
}
func l() {
	a = "O"
	print(a)
}
