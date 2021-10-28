package main

var sa = "G"

func main() {
	nn() //G
	mm() //0
	nn() //0
}
func nn() {
	print(sa)
}

func mm() {
	sa = "0"
	print(sa)
}
