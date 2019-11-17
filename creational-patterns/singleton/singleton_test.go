package singleton

import "testing"

import _ "gopkg.in/check.v1"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), but get nil")
	}

	expectedCounter := counter1

	currentCounter := counter1.AddOne()

	if currentCounter != 1 {
		t.Errorf("after calling for the first time to count, the count must be 1 but it is %d\n", currentCounter)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("expected same instance in counter2 but it got a different instance")
	}

	currentCounter = counter2.AddOne()
	if currentCounter != 2 {
		t.Errorf("after calling 'AddOne()' using the second counter, the current counter must be 2 but was %d\n", currentCounter)
	}
}
