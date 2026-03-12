package main

import (
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
	pin     string
}

func (b *BankAccount) Deposit(amount float64) {

	if amount > 0 {
		b.balance += amount
		fmt.Println("Deposit successful")
	}
}

func (b *BankAccount) Withdraw(amount float64, enteredPin string) {

	if enteredPin != b.pin {
		fmt.Println("Incorrect PIN")
		return
	}

	if amount <= b.balance {
		b.balance -= amount
		fmt.Println("Withdrawal successful")
	} else {
		fmt.Println("Insufficient balance")
	}
}

func (b BankAccount) CheckBalance(enteredPin string) {

	if enteredPin != b.pin {
		fmt.Println("Incorrect PIN")
		return
	}

	fmt.Println("Current Balance:", b.balance)
}

func day9() {

	var owner string
	var balance float64
	var pin string

	fmt.Print("Enter account holder name: ")
	fmt.Scanln(&owner)

	fmt.Print("Enter initial balance: ")
	fmt.Scanln(&balance)

	fmt.Print("Set PIN: ")
	fmt.Scanln(&pin)

	account := BankAccount{
		owner:   owner,
		balance: balance,
		pin:     pin,
	}

	for {

		fmt.Println("\n1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			var amount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scanln(&amount)
			account.Deposit(amount)

		case 2:
			var amount float64
			var enteredPin string
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scanln(&amount)
			fmt.Print("Enter PIN: ")
			fmt.Scanln(&enteredPin)
			account.Withdraw(amount, enteredPin)

		case 3:
			var enteredPin string
			fmt.Print("Enter PIN: ")
			fmt.Scanln(&enteredPin)
			account.CheckBalance(enteredPin)

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
