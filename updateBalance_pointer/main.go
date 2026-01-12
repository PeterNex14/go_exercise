package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(c *customer, transact transaction) error {
	if transact.transactionType == transactionDeposit {
		c.balance += transact.amount
	} else if transact.transactionType == transactionWithdrawal {
		if c.balance < transact.amount {
			return errors.New("insufficient funds")
		}
		c.balance -= transact.amount
	} else {
		return errors.New("unknown transaction type")
	}

	return nil
}
