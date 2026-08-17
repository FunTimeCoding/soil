package expected

import (
	"expected.test/assert"
	"expected.test/constant"
	"fmt"
	"math"
	"testing"
)

type Item struct {
	Name string
}

func New(name string) *Item {
	return &Item{Name: name}
}

func Direct(
	t *testing.T,
	actual string,
) {
	assert.String(t, "name", actual)
}

func Nested(
	t *testing.T,
	actual *Item,
) {
	assert.Item(t, New("name"), actual)
}

func Helper(
	t *testing.T,
	actual *Item,
) {
	assertItem(t, New("name"), actual)
}

func assertItem(
	t *testing.T,
	expected *Item,
	actual *Item,
) {
	t.Helper()
	assert.Item(t, expected, actual)
}

func Flagged() {
	fmt.Println("name") // want `string literal "name" has constant constant.Name`
}

func Tautology(
	t *testing.T,
	actual string,
) {
	assert.String(
		t,
		constant.Name,
		actual,
	) // want `constant constant.Name in expected value should be a literal`
}

func NestedTautology(
	t *testing.T,
	actual *Item,
) {
	assert.Item(
		t,
		New(constant.Name),
		actual,
	) // want `constant constant.Name in expected value should be a literal`
}

func UniverseClean(
	t *testing.T,
	actual any,
) {
	assert.Item(t, true, actual)
}

func StandardLibraryClean(
	t *testing.T,
	actual any,
) {
	assert.Item(t, math.Pi, actual)
}

func Matches(
	s string,
	list []string,
) bool {
	return len(list) > 0 && list[0] == s
}

func BooleanConstantClean(t *testing.T) {
	assert.True(t, Matches(constant.Name, nil))
}

func BooleanLiteralFlagged(t *testing.T) {
	assert.True(
		t,
		Matches("name", nil),
	) // want `string literal "name" has constant constant.Name`
}

func lookup(name string) map[string]float64 {
	return map[string]float64{name: 0}
}

func IndexClean(
	t *testing.T,
	actual float64,
	table map[string]float64,
) {
	assert.Item(t, table[constant.Name], actual)
}

func SliceClean(
	t *testing.T,
	actual string,
	text string,
) {
	assert.String(t, text[constant.Start:], actual)
}

func IndexContainerFlagged(
	t *testing.T,
	actual float64,
) {
	assert.Item(
		t,
		lookup(constant.Name)[constant.Name],
		actual,
	) // want `constant constant.Name in expected value should be a literal`
}
