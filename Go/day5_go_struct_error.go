package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name   string
	age    int
	course string
}

type Book struct {
	title  string
	author string
	price  float64
}

func day5() {

	fmt.Println("----- Struct Introduction -----")

	s1 := Student{
		name:   "Niraj",
		age:    22,
		course: "CSE",
	}

	fmt.Println("Student Name:", s1.name)
	fmt.Println("Age:", s1.age)
	fmt.Println("Course:", s1.course)

	fmt.Println("\n----- Struct Practice -----")

	book1 := Book{
		title:  "Learning Go",
		author: "John Doe",
		price:  599.99,
	}

	fmt.Println("Book Title:", book1.title)
	fmt.Println("Author:", book1.author)
	fmt.Println("Price:", book1.price)

	fmt.Println("\n----- Error Handling Example -----")

	var input string

	fmt.Print("Enter a number: ")
	fmt.Scan(&input)

	number, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error: Invalid number entered")
		return
	}

	fmt.Println("You entered:", number)
}
