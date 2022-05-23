package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "tea",
		Price: 1,
		SKU: "ygduyg-fffffff",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
