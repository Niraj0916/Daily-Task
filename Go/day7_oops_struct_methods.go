package main

import "fmt"

type Calculator struct{}

func (c Calculator) Add(a, b float64) float64 {
	return a + b
}

func (c Calculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c Calculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c Calculator) Divide(a, b float64) string {
	if b == 0 {
		return "Error: Cannot divide by zero"
	}
	return fmt.Sprintf("%.2f", a/b)
}

type book struct {
	Title  string
	Author string
	Pages  int
}

func (b book) displayInfo() string {
	return fmt.Sprintf("'%s' by %s, %d pages", b.Title, b.Author, b.Pages)
}

func day7() {

	// Calculator Example
	calc := Calculator{}

	fmt.Printf("10 + 5 = %.2f\n", calc.Add(10, 5))
	fmt.Printf("10 - 5 = %.2f\n", calc.Subtract(10, 5))
	fmt.Printf("10 * 5 = %.2f\n", calc.Multiply(10, 5))
	fmt.Printf("10 / 5 = %s\n", calc.Divide(10, 5))
	fmt.Printf("10 / 0 = %s\n\n", calc.Divide(10, 0))

	// Book Example
	book1 := book{
		Title:  "Python Crash Course",
		Author: "Eric Matthes",
		Pages:  544,
	}

	book2 := book{
		Title:  "Clean Code",
		Author: "Robert Martin",
		Pages:  464,
	}

	fmt.Println(book1.displayInfo())
	fmt.Println(book2.displayInfo())
}
