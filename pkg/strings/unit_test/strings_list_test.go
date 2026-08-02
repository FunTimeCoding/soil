package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/unit_test/strings_tester"
	"testing"
)

func TestCompare(t *testing.T) {
	assertCompare(
		t,
		strings_tester.Compare{},
		[]string{},
		[]string{},
	)
	assertCompare(
		t,
		strings_tester.Compare{Add: []string{"Alfa"}},
		[]string{},
		[]string{constant.UpperAlfa},
	)
	assertCompare(
		t,
		strings_tester.Compare{Remove: []string{"Alfa"}},
		[]string{constant.UpperAlfa},
		[]string{},
	)
	assertCompare(
		t,
		strings_tester.Compare{Stay: []string{"Alfa"}},
		[]string{constant.UpperAlfa},
		[]string{constant.UpperAlfa},
	)
}

func assertCompare(
	t *testing.T,
	expected strings_tester.Compare,
	past []string,
	now []string,
) {
	t.Helper()
	add, remove, stay := strings.Compare(past, now)
	assert.Strings(t, notNil(expected.Add), add)
	assert.Strings(t, notNil(expected.Remove), remove)
	assert.Strings(t, notNil(expected.Stay), stay)
}

func notNil(s []string) []string {
	if s == nil {
		return []string{}
	}

	return s
}

func TestDeleteDuplicates(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		strings.DeleteDuplicates(
			[]string{
				constant.UpperAlfa,
				constant.UpperAlfa,
				constant.UpperBravo,
			},
		),
	)
}

func TestDeleteEmpty(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		strings.DeleteEmpty(
			[]string{
				"",
				constant.UpperAlfa,
				"",
				constant.UpperBravo,
				"",
			},
		),
	)
}

func TestFirstHalf(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		strings.FirstHalf(
			[]string{
				constant.UpperAlfa,
				constant.UpperBravo,
				constant.UpperCharlie,
				constant.UpperDelta,
			},
		),
	)
}

func TestLinesAfter(t *testing.T) {
	assert.Strings(
		t,
		[]string{"c"},
		strings.LinesAfter([]string{"a", "b", "c"}, "b"),
	)
}

func TestRemoveFromList(t *testing.T) {
	assert.Any(
		t,
		[]string{"Bravo", "Charlie"},
		strings.RemoveFromList(
			[]string{
				constant.UpperAlfa,
				constant.UpperAlfa,
				constant.UpperBravo,
				constant.UpperCharlie,
			},
			[]string{constant.UpperAlfa},
		),
	)
	assert.Any(
		t,
		[]string{"Charlie"},
		strings.RemoveFromList(
			[]string{
				constant.UpperAlfa,
				constant.UpperAlfa,
				constant.UpperBravo,
				constant.UpperCharlie,
			},
			[]string{constant.UpperAlfa, constant.UpperBravo},
		),
	)
}

func TestReverse(t *testing.T) {
	reversed := []string{
		constant.UpperCharlie,
		constant.UpperBravo,
		constant.UpperAlfa,
	}
	strings.Reverse(reversed)
	assert.Any(
		t,
		[]string{"Alfa", "Bravo", "Charlie"},
		reversed,
	)
}

func TestSecondHalf(t *testing.T) {
	assert.Any(
		t,
		[]string{"Charlie", "Delta"},
		strings.SecondHalf(
			[]string{
				constant.UpperAlfa,
				constant.UpperBravo,
				constant.UpperCharlie,
				constant.UpperDelta,
			},
		),
	)
}

func TestSort(t *testing.T) {
	preSorted := []string{
		constant.UpperAlfa,
		constant.UpperBravo,
		constant.UpperCharlie,
	}
	strings.Sort(preSorted, true)
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo", "Charlie"},
		preSorted,
	)
	ascending := []string{
		constant.UpperBravo,
		constant.UpperAlfa,
		constant.UpperCharlie,
	}
	strings.Sort(ascending, true)
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo", "Charlie"},
		ascending,
	)
	descending := []string{
		constant.UpperBravo,
		constant.UpperAlfa,
		constant.UpperCharlie,
	}
	strings.Sort(descending, false)
	assert.Strings(
		t,
		[]string{"Charlie", "Bravo", "Alfa"},
		descending,
	)
}

func TestSwap(t *testing.T) {
	actual := []string{
		constant.UpperAlfa,
		constant.UpperBravo,
		constant.UpperCharlie,
	}
	strings.Swap(actual, 1, 2)
	assert.Any(
		t,
		[]string{"Alfa", "Charlie", "Bravo"},
		actual,
	)
}
