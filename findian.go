package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')

	// Trim spaces and convert to lowercase for case-insensitive comparison
	input = strings.TrimSpace(strings.ToLower(input))
	// Check conditions: starts with 'i', contains 'a', and ends with 'n'
	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
