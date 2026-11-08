package target

import "testing"

func (y *Shelf) label() string {
	return y.product.Name
}

func TestLabel(t *testing.T) {
	s := &Shelf{product: &Product{Name: "a"}}

	if s.label() != "a" {
		t.Fatal("label")
	}
}
