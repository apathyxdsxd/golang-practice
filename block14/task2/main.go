package main

import "fmt"

type BankAccount struct {
	balance float64
}

func (a *BankAccount) Dep(amount float64) {
	a.balance += amount
}

func (a *BankAccount) Withdraw(amount float64) error {
	if amount > a.balance {
		return fmt.Errorf("недостаточно средств на счёте")
	}
	a.balance -= amount
	return nil
}

func (a *BankAccount) CheckBalance() float64 {
	return a.balance
}

func main() {
	account := BankAccount{balance: 1000.0}

	fmt.Println("Исходный баланс:", account.CheckBalance())

	account.Dep(500.0)
	fmt.Println("Баланс после пополнения:", account.CheckBalance())

	err := account.Withdraw(200.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("баланс после снятия:", account.CheckBalance())
	}
}
