package main

import "fmt"

type Animal struct {
	Name string
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

func (d Dog) Speak() string {
	return "Woof"
}

func (c Cat) Speak() string {
	return "Meow"
}

type Speaker interface {
	Speak() string
}

func day8() {

	dog := Dog{Animal{Name: "Buddy"}}
	cat := Cat{Animal{Name: "Whiskers"}}

	animals := []Speaker{dog, cat}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

}
