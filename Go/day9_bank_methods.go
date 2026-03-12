package main

import "fmt"

type BankAccount struct {
	owner   string
	balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.balance += amount
		fmt.Println("Deposit successful")
	}
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount <= b.balance {
		b.balance -= amount
		fmt.Println("Withdrawal successful")
	} else {
		fmt.Println("Insufficient funds")
	}
}

func (b BankAccount) GetBalance() float64 {
	return b.balance
}

func day9() {

	account := BankAccount{
		owner:   "Niraj",
		balance: 1000,
	}

	account.Deposit(500)
	account.Withdraw(300)

	fmt.Println("Balance:", account.GetBalance())

}
