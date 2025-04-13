package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name of the text file: ")
	fileNameInput, _ := reader.ReadString('\n')

	fileName := strings.TrimSpace(fileNameInput)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print("Error opening file: ", err)
		return
	}
	defer file.Close()

	var names []Name
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			fmt.Println("Skipping malformed line:", line)
			continue
		}

		// Truncate to 20 characters if needed
		fname := parts[0]
		lname := parts[1]
		if len(fname) > 20 {
			fname = fname[:20]
		}
		if len(lname) > 20 {
			lname = lname[:20]
		}

		names = append(names, Name{
			Fname: fname,
			Lname: lname,
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print all names
	fmt.Println("\nNames found in file:")
	for _, name := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", name.Fname, name.Lname)
	}
}
