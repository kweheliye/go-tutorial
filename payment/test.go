package main

import (
	"fmt"
	"time"
)

type ProcessPayment interface {
	processPayment(money float64) error
}

type CreditPayment struct {
	cardNo     int
	expiryDate time.Time
}

type WalletPayment struct {
	walletId string
	balance  float64
}

func (c CreditPayment) processPayment(amount float64) error {

	if time.Now().After(c.expiryDate) {
		return fmt.Errorf("credit card expired")
	}
	fmt.Println("Credit Payment Successful")
	return nil
}

func (w WalletPayment) processPayment(amount float64) error {
	if w.balance < amount {
		return fmt.Errorf("not enough balance in wallet")
	}
	w.balance -= amount
	fmt.Println("Wallet Payment Successful")
	return nil
}

func main() {

	credit := CreditPayment{cardNo: 12234, expiryDate: time.Now().AddDate(1, 0, 0)}

	if err := credit.processPayment(100); err != nil {
		fmt.Println("Error:", err)
	}

}
