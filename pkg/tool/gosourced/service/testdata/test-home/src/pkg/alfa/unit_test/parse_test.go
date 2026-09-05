package unit_test

import (
	"example/pkg/alfa"
	"example/pkg/alfa/unit_test/alfa_tester"
	"testing"
)

func TestParse(t *testing.T) {
	if alfa.Parse(alfa_tester.Sample()) != "sample" {
		t.Fatal("parse")
	}
}
