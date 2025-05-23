package main

import "errors"

type Account struct {
	balance float64
	owner   string
}

func NewAccount(owner string) *Account {
	return &Account{owner: owner, balance: 0.0}
}

func (a *Account) SetBalance(value float64) error {
	if value < 0 {
		return errors.New("negative balance not allowed")
	}
	a.balance = value
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}
func (a *Account) Deposit(value float64) error {
	if value < 0 {
		return errors.New("negative balance not allowed")
	}
	a.balance += value
	return nil
}
func (a *Account) Withdraw(value float64) error {
	if value > a.balance || value < 0 {
		return errors.New("balance not allowed")
	}
	a.balance -= value
	return nil
}
