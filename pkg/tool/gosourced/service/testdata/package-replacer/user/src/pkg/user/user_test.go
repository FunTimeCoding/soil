package user

import (
	"other.test/lib/shape"
	"testing"
)

func TestWidth(t *testing.T) {
	if Width() != shape.New().Width {
		t.Fail()
	}
}
