package main

import "fmt"

func day4_arrays() {
	// Arrays
	var fruits [3]string = [3]string{"Apple", "Banana", "Mango"}
	fmt.Println("Array:", fruits)
	fmt.Println("First fruit:", fruits[0])

	// Slices
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice:", numbers)
	fmt.Println("Slice length:", len(numbers))

	// Append to slice
	numbers = append(numbers, 60)
	fmt.Println("After append:", numbers)

	// Slice of a slice
	fmt.Println("Slice[1:3]:", numbers[1:3])
}
