package main

import "fmt"

func day3() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("The number is Even")
	} else {
		fmt.Println("The number is Odd")
	}

	var a, b int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Maximum number is:", a)
	} else {
		fmt.Println("Maximum number is:", b)
	}
}
