package main

import "fmt"

func main() {

	// Multiplication Table

	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	fmt.Println("Multiplication Table:")
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}

	// Factorial

	var n int
	fmt.Print("Enter number to find factorial: ")
	fmt.Scan(&n)

	factorial := 1

	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Println("Factorial is:", factorial)

	// Fibonacci Series

	var terms int
	fmt.Print("Enter number of Fibonacci terms: ")
	fmt.Scan(&terms)

	a := 0
	b := 1

	fmt.Println("Fibonacci Series:")

	for i := 0; i < terms; i++ {
		fmt.Print(a, " ")
		next := a + b
		a = b
		b = next
	}
}
