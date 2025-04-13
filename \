package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	nameInput, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameInput)
	fmt.Print("Enter address: ")
	addressInput, _ := reader.ReadString('\n')
	address := strings.TrimSpace(addressInput)
	data := map[string]string{
		"name":    name,
		"address": address,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to json: ", err)
		return
	}
	fmt.Print("Printing json: ")
	fmt.Print(string(jsonData))
}
