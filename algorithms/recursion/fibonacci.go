package main

import "fmt"

func fibonacci(num int) int {
	if num == 2 || num == 1 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func fibo_seq(num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(fibonacci(i))
	}
}

func main() {
	fibo_seq(9)
}
