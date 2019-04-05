// -*- go-use-test-args: "-race"; -*-

package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtCloner()

	if shirtCache == nil {
		t.Fatal("Received cache is nil\n")
	}

	item1, err := shirtCache.GetClone(White)

	if err != nil {
		t.Error(err)
	}

	if item1 == whitePrototype {
		t.Error("item1 can not be equal to the white prototype")
	}

	shirt1, ok := item1.(*Shirt)

	if !ok {
		t.Fatal("item1 is not *Shirt type")
	}

	shirt1.SKU = "abbcc"

	item2, err := shirtCache.GetClone(White)

	if err != nil {
		t.Error(err)
	}

	shirt2, ok := item2.(*Shirt)

	if !ok {
		t.Fatal("item2 is not *Shirt type")
	}

	if shirt2 == shirt1 {
		t.Fatal("shirt1 cannot be equal to shirt2")
	}

	if shirt2.SKU == shirt1.SKU {
		t.Fatal("SKU of shirt1 and shirt2 should be different")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())

	t.Logf("LOG: memory location of shirt1 (%p) != shirt2 (%p)", shirt1, shirt2)
}
