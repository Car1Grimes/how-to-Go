package main

import (
	"fmt"
)

func BubbleSort(numbers []int) {
	length := len(numbers)
	for outer := 0; outer < length-1; outer++ {
		for inner := 0; inner < length-1-outer; inner++ {
			Swap(numbers, inner)
		}
	}
}

func Swap(numbers []int, index int) {
	if numbers[index] > numbers[index+1] {
		temp := numbers[index]
		numbers[index] = numbers[index+1]
		numbers[index+1] = temp
		return
	}
}

func main() {
	var num int
	var sample []int
	fmt.Print("Enter a single number at a time for sorting (10 numbers max, any NaN key to quit): ")
	for i := 0; i <= 9; i++ {
		_, err := fmt.Scanln(&num)
		if err != nil {
			break
		}
		sample = append(sample, num)
		fmt.Printf("%d numbers left\n", 9-i)
	}
	fmt.Println("Your unsorted array: ")
	for _, value := range sample {
		fmt.Printf("%d ", value)
	}
	BubbleSort(sample)
	fmt.Println("\nYour sorted array: ")
	for _, value := range sample {
		fmt.Printf("%d ", value)
	}
}
