package factory

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(CASH)

	if err != nil {
		t.Fatal("a payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)

	if !strings.Contains(msg, "paid using cash") {
		t.Error("the cash payment method message wasn't correct")
	}

	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DEBIT_CARD)

	if err != nil {
		t.Error("a payment method of type 'DebitCard' must exist")
	}

	msg := payment.Pay(22.20)

	if !strings.Contains(msg, "paid using debit card") {
		t.Error("the debit card payment method message wasn't correct")
	}

	t.Log("LOG:", msg)
}

func TestGetPaymentNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(UNSUPPORTED)

	if err == nil {
		t.Error("a payment method with ID UNSUPPORTED must return an error")
	}

	t.Log("LOG:", err)
}
