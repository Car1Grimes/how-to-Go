package main

import (
	"fmt"
	"time"
)

var counter = 1

func increment() {
	for i := 1; i < 1000; i++ {
		counter++
	}
}
func main() {
	increment()
	increment()
	time.Sleep(time.Second)

	fmt.Println("final counter value: ", counter)
}

//Explanation:
//When two increment() run concurrently, they might read the same value (say 5), both increment it to 6 and both write it back (write 6,not 7)
//This means some increments are lost
