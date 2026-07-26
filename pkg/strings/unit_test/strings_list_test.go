package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestCompare(t *testing.T) {
	assertCompare(
		t,
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
	)
	// Add
	assertCompare(
		t,
		[]string{constant.UpperAlfa},
		[]string{},
		[]string{},
		[]string{},
		[]string{constant.UpperAlfa},
	)
	// Remove
	assertCompare(
		t,
		[]string{},
		[]string{constant.UpperAlfa},
		[]string{},
		[]string{constant.UpperAlfa},
		[]string{},
	)
	// Stay
	assertCompare(
		t,
		[]string{},
		[]string{},
		[]string{constant.UpperAlfa},
		[]string{constant.UpperAlfa},
		[]string{constant.UpperAlfa},
	)
}

func assertCompare(
	t *testing.T,
	expectAdd []string,
	expectRemove []string,
	expectStay []string,
	past []string,
	now []string,
) {
	t.Helper()
	add, remove, stay := strings.Compare(past, now)
	assert.Any(t, expectAdd, add)
	assert.Any(t, expectRemove, remove)
	assert.Any(t, expectStay, stay)
}

func TestDeleteDuplicates(t *testing.T) {
	assert.Any(
		t,
		[]string{constant.UpperAlfa, constant.UpperBravo},
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
		strings.DeleteEmpty([]string{"", constant.UpperAlfa, "", constant.UpperBravo, ""}),
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
		[]string{constant.UpperBravo, constant.UpperCharlie},
		strings.RemoveFromList(
			[]string{constant.UpperAlfa, constant.UpperAlfa, constant.UpperBravo, constant.UpperCharlie},
			[]string{constant.UpperAlfa},
		),
	)
	assert.Any(
		t,
		[]string{constant.UpperCharlie},
		strings.RemoveFromList(
			[]string{constant.UpperAlfa, constant.UpperAlfa, constant.UpperBravo, constant.UpperCharlie},
			[]string{constant.UpperAlfa, constant.UpperBravo},
		),
	)
}

func TestReverse(t *testing.T) {
	reversed := []string{constant.UpperCharlie, constant.UpperBravo, constant.UpperAlfa}
	strings.Reverse(reversed)
	assert.Any(
		t,
		[]string{constant.UpperAlfa, constant.UpperBravo, constant.UpperCharlie},
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
	preSorted := []string{constant.UpperAlfa, constant.UpperBravo, constant.UpperCharlie}
	strings.Sort(preSorted, true)
	assert.Strings(
		t,
		[]string{constant.UpperAlfa, constant.UpperBravo, constant.UpperCharlie},
		preSorted,
	)
	ascending := []string{constant.UpperBravo, constant.UpperAlfa, constant.UpperCharlie}
	strings.Sort(ascending, true)
	assert.Strings(
		t,
		[]string{constant.UpperAlfa, constant.UpperBravo, constant.UpperCharlie},
		ascending,
	)
	descending := []string{constant.UpperBravo, constant.UpperAlfa, constant.UpperCharlie}
	strings.Sort(descending, false)
	assert.Strings(
		t,
		[]string{constant.UpperCharlie, constant.UpperBravo, constant.UpperAlfa},
		descending,
	)
}

func TestSwap(t *testing.T) {
	actual := []string{constant.UpperAlfa, constant.UpperBravo, constant.UpperCharlie}
	strings.Swap(actual, 1, 2)
	assert.Any(
		t,
		[]string{constant.UpperAlfa, constant.UpperCharlie, constant.UpperBravo},
		actual,
	)
}
