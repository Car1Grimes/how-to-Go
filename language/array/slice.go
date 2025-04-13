package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers := make([]int, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Please enter an integer or press X to quit: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if strings.ToUpper(input) == "X" {
			fmt.Print("Exiting...")
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input. Please enter an integer or press X to quit: ")
			continue
		}
		numbers = append(numbers, num)
		sort.Ints(numbers)

		fmt.Println("Sorted numbers: ", numbers)
	}
}
