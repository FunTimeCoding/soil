package math

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestExceededByPercent(t *testing.T) {
	// No time change
	assertExceededByPercent(
		t,
		100,
		100,
		100,
		10,
		false,
	)
	// 9% exceeded
	assertExceededByPercent(
		t,
		100,
		109,
		100,
		10,
		false,
	)
	// 10% exceeded
	assertExceededByPercent(
		t,
		100,
		110,
		100,
		10,
		true,
	)
	// 11% exceeded
	assertExceededByPercent(
		t,
		100,
		111,
		100,
		10,
		true,
	)
	// 19% exceeded
	assertExceededByPercent(
		t,
		100,
		119,
		100,
		10,
		true,
	)
	// 20% exceeded
	assertExceededByPercent(
		t,
		100,
		120,
		100,
		10,
		true,
	)
	// 21% exceeded
	assertExceededByPercent(
		t,
		100,
		121,
		100,
		10,
		true,
	)
	// 9% exceeded
	assertExceededByPercent(
		t,
		110,
		119,
		100,
		10,
		false,
	)
	// 10% exceeded
	assertExceededByPercent(
		t,
		110,
		120,
		100,
		10,
		true,
	)
	// 11% exceeded
	assertExceededByPercent(
		t,
		110,
		121,
		100,
		10,
		true,
	)
	assertExceededByPercent(
		t,
		661399.2-1,
		661399.2,
		601272,
		10,
		true,
	)
}

func assertExceededByPercent(
	t *testing.T,
	past float64,
	present float64,
	hundredPercent float64,
	percent float64,
	expect bool,
) {
	t.Helper()
	assert.Boolean(
		t,
		expect,
		ExceededByPercent(past, present, hundredPercent, percent),
	)
}
