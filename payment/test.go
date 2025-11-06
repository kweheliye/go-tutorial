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

func processPayment(p ProcessPayment, amount float64) {
	if err := p.processPayment(amount); err != nil {
		fmt.Println("Payment failed:", err)
	} else {
		fmt.Println("Payment Processed Successful")
	}
}

func PaymentFactory(paymentType string, data map[string]interface{}) ProcessPayment {
	switch paymentType {
	case "credit":
		return CreditPayment{
			cardNo:     data["cardNo"].(int),
			expiryDate: data["expiryDate"].(time.Time),
		}
	case "wallet":
		return WalletPayment{
			walletId: data["walletId"].(string),
			balance:  data["balance"].(float64),
		}
	}
	return nil
}

func main() {

	credit := PaymentFactory("credit", map[string]interface{}{
		"cardNo":     123456789,
		"expiryDate": time.Now().AddDate(1, 0, 0),
	})

	wallet := PaymentFactory("wallet", map[string]any{
		"walletId": "12345",
		"balance":  1000.00,
	})

	processPayment(credit, 100)
	processPayment(wallet, 100)

}
