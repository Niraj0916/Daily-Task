package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a float64, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Cannot divide by zero")
		return 0
	}
	return a / b
}

func day4_functions() {
	var a, b int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("Add:", add(a, b))
	fmt.Println("Subtract:", subtract(a, b))
	fmt.Println("Multiply:", multiply(a, b))
	fmt.Println("Divide:", divide(float64(a), float64(b)))
}
