package main

var ax = "G"

func main() {
	nx() //G
	mx() //0
	nx() //G
}
func nx() {
	print(ax)
}

func mx() {
	ax := "0"
	print(ax)
}
