package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := NewShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Not implemented yet")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}
	if item1 == whitePrototype {
		t.Fatal("item1 connot be equal to the white prototype")
	}
	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type aassertion for shirt1 couldn't be done successfully")
	}
	shirt1.SKU = "abbcc"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}
	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type aassertion for shirt2 couldn't be done successfully")
	}
	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}
	if shirt1 == shirt2 {
		t.Error("shirt1 cannot be equal to shirt2")
	}
	t.Logf("LOG - shirt1: %s", shirt1.GetInfo())
	t.Logf("LOG - shirt2: %s", shirt2.GetInfo())
}
