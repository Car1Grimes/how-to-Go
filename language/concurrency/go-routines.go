package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Enter a series of integers (space-separated):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Parse input
	fields := strings.Fields(input)
	var nums []int
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			fmt.Println("Invalid input. Please enter only integers.")
			return
		}
		nums = append(nums, num)
	}

	if len(nums) == 0 {
		fmt.Println("No numbers entered.")
		return
	}

	// Partition array into 4 parts
	partitions := partition(nums, 4)

	var wg sync.WaitGroup
	sortedChan := make(chan []int, 4)

	// Sort each partition in a goroutine
	for _, part := range partitions {
		wg.Add(1)
		go func(p []int) {
			defer wg.Done()
			fmt.Println("Sorting subarray:", p)
			sort.Ints(p)
			sortedChan <- p
		}(part)
	}

	wg.Wait()
	close(sortedChan)

	// Collect sorted subarrays
	var sortedSubarrays [][]int
	for sorted := range sortedChan {
		sortedSubarrays = append(sortedSubarrays, sorted)
	}

	// Merge all sorted subarrays
	finalSorted := mergeSortedArrays(sortedSubarrays)
	fmt.Println("Final sorted array:", finalSorted)
}

// partition splits the array into n roughly equal parts
func partition(arr []int, n int) [][]int {
	var result [][]int
	chunkSize := (len(arr) + n - 1) / n // ceiling division
	for i := 0; i < len(arr); i += chunkSize {
		end := i + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}
	return result
}

// mergeSortedArrays merges multiple sorted slices into one sorted slice
func mergeSortedArrays(arrays [][]int) []int {
	if len(arrays) == 0 {
		return nil
	}
	result := arrays[0]
	for i := 1; i < len(arrays); i++ {
		result = mergeTwoSorted(result, arrays[i])
	}
	return result
}

// mergeTwoSorted merges two sorted slices
func mergeTwoSorted(a, b []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}
