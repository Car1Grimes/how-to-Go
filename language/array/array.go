package main

import "fmt"

func main() {
	var	arr = [5]int{1,2,3,4,5}
	for i, v := range arr {
		fmt.Printf("index: %d, value: %d", i, v)
		fmt.Println("")
	}
	return
}
