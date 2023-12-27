package main

import (
	"fmt"
)

type Memento2 struct {
	mementoBalance int
}

type BankAccount2 struct {
	balance int
	changes []*Memento2
	counter int
}

func NewBankAccount2(balance int) *BankAccount2 {
	ba := &BankAccount2{balance: balance}
	mo := &Memento2{mementoBalance: ba.balance}
	ba.changes = append(ba.changes, mo)
	return ba
}

func (ba *BankAccount2) Deposit2(amount int) *Memento2 {
	ba.balance += amount
	mo := &Memento2{ba.balance}
	ba.changes = append(ba.changes, mo)
	ba.counter++
	fmt.Println("Deposited", amount,
		", balance is now", ba.balance)
	return mo

}

func (ba *BankAccount2) Undo() {
	if ba.counter > 0 {
		ba.counter--
		memento := ba.changes[ba.counter]
		ba.balance = memento.mementoBalance
	}else{
		fmt.Println("You can not undo: ", ba.balance)
	}
}

func main() {
	bankAccount := NewBankAccount2(100)
	bankAccount.Deposit2(20)
	bankAccount.Deposit2(200)
	bankAccount.Undo()
	// bankAccount.Undo()
	// bankAccount.Undo()
	fmt.Println("Back to: ", bankAccount.balance,
		", counter:", bankAccount.counter)

}
