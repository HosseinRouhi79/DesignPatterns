package main

import "fmt"

type Memento struct {
	mementoBalance int
}

type BankAccount struct {
	balance int
}

func NewBankAccount(balance int) (*BankAccount, *Memento) {
	ba := &BankAccount{balance: balance}
	mo := &Memento{mementoBalance: ba.balance}

	return ba, mo
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	return &Memento{b.balance}
}

func (b *BankAccount) Restore(m *Memento) {
	b.balance = m.mementoBalance
}

func main() {
	bankAccount, mInit := NewBankAccount(100)
	m1 := bankAccount.Deposit(20)
	m2 := bankAccount.Deposit(200)
	fmt.Println(bankAccount)

	fmt.Println(mInit)
	bankAccount.Restore(m1)
	fmt.Println(bankAccount)
	bankAccount.Restore(m2)
	fmt.Println(bankAccount)

}
