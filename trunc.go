package main

import "fmt"

func main() {
	var num float64
	fmt.Print("please enter a floating number: ")
	_,err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Wrong float format of the entered number")
		return
	}
	truncate := int(num)
	fmt.Println(truncate)
}
