package main

import "fmt"

func f() {
	var x int = 3
	fmt.Println(x)
}

func g() {
	var x int = 4
	fmt.Println(x)
}


func main() {
	f()
	g()
}
