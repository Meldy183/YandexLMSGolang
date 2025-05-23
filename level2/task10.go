package main

import (
	"errors"
)

var amErr = errors.New("amount is incorrect")
var balErr = errors.New("balance is incorrect")
var Balance = 0.0

func makeBalanceZero() {
	Balance = 0.0
}
func topUpBalance(amount float64) error {
	if r := recover(); r != nil {
		return amErr
	}
	if amount <= 0 {
		return amErr
	}
	Balance += amount
	return nil
}

func chargeFromBalance(amount float64) error {
	if r := recover(); r != nil {
		return amErr
	}
	if amount <= 0 {
		return amErr
	}
	Balance -= amount
	return nil
}

func TopUpAndGetBalance(amount float64) (float64, error) {
	err := topUpBalance(amount)
	if err != nil {
		return 0, err
	}
	if Balance < 0 {
		return 0, balErr
	}
	return Balance, nil
}

func ChargeFromAndGetBalance(amount float64) (float64, error) {
	err := chargeFromBalance(amount)
	if err != nil {
		return 0, err
	}
	if Balance < 0 {
		return 0, balErr
	}
	return Balance, nil
}
