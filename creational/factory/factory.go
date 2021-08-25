package factory

import (
	"errors"
	"fmt"
)

const (
	Cash      = 1
	DebitCard = 2
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m int) (PaymentMethod, error) {

	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(CreditCardPM), nil
	default:
		return nil, errors.New("no implemented yet")
	}
}

type CashPM struct{}
type DebitCardPM struct{}
type CreditCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}

func (d *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card (new)\n", amount)
}
