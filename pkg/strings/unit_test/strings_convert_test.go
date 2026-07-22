package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestAnyToSlice(t *testing.T) {
	assert.Strings(t, []string{"Alfa"}, strings.AnyToSlice(upper.Alfa))
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		strings.AnyToSlice([]string{upper.Alfa, upper.Bravo}),
	)
}

func TestMustToFloat(t *testing.T) {
	assert.Float(t, 1, strings.MustToFloat("1.0"))
	assert.Float(t, 1, strings.MustToFloat(" 1.0"))
}

func TestMustToInteger64(t *testing.T) {
	assert.Integer64(t, 1, strings.MustToInteger64(strings.One))
	assert.Integer64(t, 1, strings.MustToInteger64(" 1"))
}

func TestMustToInteger(t *testing.T) {
	assert.Integer(t, 1, strings.MustToInteger(strings.One))
	assert.Integer(t, 1, strings.MustToInteger(" 1"))
}

func TestMustToIntegers64(t *testing.T) {
	assert.Any(t, []int64{}, strings.MustToIntegers64([]string{}))
	assert.Any(t, []int64{5}, strings.MustToIntegers64([]string{"5"}))
	assert.Any(t, []int64{1, 2}, strings.MustToIntegers64([]string{"1", "2"}))
}

func TestMustToIntegers(t *testing.T) {
	assert.Any(t, []int{}, strings.MustToIntegers([]string{}))
	assert.Any(t, []int{5}, strings.MustToIntegers([]string{"5"}))
	assert.Any(t, []int{1, 2}, strings.MustToIntegers([]string{"1", "2"}))
}

func TestMustToUnsignedInteger(t *testing.T) {
	assert.Unsigned(t, 1, strings.MustToUnsignedInteger(strings.One))
	assert.Unsigned(t, 1, strings.MustToUnsignedInteger(" 1"))
}

func TestToBoolean(t *testing.T) {
	assert.True(t, strings.ToBoolean(strings.BooleanTrue))
	assert.False(t, strings.ToBoolean(strings.BooleanFalse))
}

func TestToFloat(t *testing.T) {
	assert.Float(t, 0.5, strings.ToFloat("0.5", 0))
	assert.Float(t, 0.5, strings.ToFloat(" 0.5", 0))
}

func TestToInteger(t *testing.T) {
	assert.Integer(t, 5, strings.ToInteger(strings.Five, 0))
	assert.Integer(t, 5, strings.ToInteger(" 5", 0))
}

func TestToLower(t *testing.T) {
	assert.Strings(
		t,
		[]string{"alfa", "bravo"},
		strings.ToLower([]string{upper.Alfa, upper.Bravo}),
	)
}

func TestToMap(t *testing.T) {
	assert.Any(
		t,
		map[string]string{"A": "a", "B": "b"},
		strings.ToMap([]string{"A=a", "B=b"}, "="),
	)
}

func TestUnsignedToInteger(t *testing.T) {
	assert.Unsigned(t, 5, strings.ToUnsignedInteger(strings.Five, 0))
	assert.Unsigned(t, 5, strings.ToUnsignedInteger(" 5", 0))
}
