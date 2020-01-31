package factory

import (
	"fmt"
)

type PaymentType int

const (
	CASH PaymentType = iota + 1
	DEBIT_CARD
	UNSUPPORTED
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m PaymentType) (PaymentMethod, error) {
	switch m {
	case CASH:
		return new(CashPM), nil

	case DEBIT_CARD:
		return new(DebitCardPM), nil

	default:
		return nil, fmt.Errorf("payment method %d not recognized\n", m)
	}
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}
