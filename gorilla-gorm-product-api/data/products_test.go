package data

import "testing"

func TestChecksValidation(t*testing.T) {
	p := &Product{
		Name: "Coffe",
		Price: 1.00,
		SKU: "abs-abc-asd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}