package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/upper"
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
		[]string{upper.Alfa},
		[]string{},
		[]string{},
		[]string{},
		[]string{upper.Alfa},
	)
	// Remove
	assertCompare(
		t,
		[]string{},
		[]string{upper.Alfa},
		[]string{},
		[]string{upper.Alfa},
		[]string{},
	)
	// Stay
	assertCompare(
		t,
		[]string{},
		[]string{},
		[]string{upper.Alfa},
		[]string{upper.Alfa},
		[]string{upper.Alfa},
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
		[]string{upper.Alfa, upper.Bravo},
		strings.DeleteDuplicates(
			[]string{
				upper.Alfa,
				upper.Alfa,
				upper.Bravo,
			},
		),
	)
}

func TestDeleteEmpty(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		strings.DeleteEmpty([]string{"", upper.Alfa, "", upper.Bravo, ""}),
	)
}

func TestFirstHalf(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		strings.FirstHalf(
			[]string{
				upper.Alfa,
				upper.Bravo,
				upper.Charlie,
				upper.Delta,
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
		[]string{upper.Bravo, upper.Charlie},
		strings.RemoveFromList(
			[]string{upper.Alfa, upper.Alfa, upper.Bravo, upper.Charlie},
			[]string{upper.Alfa},
		),
	)
	assert.Any(
		t,
		[]string{upper.Charlie},
		strings.RemoveFromList(
			[]string{upper.Alfa, upper.Alfa, upper.Bravo, upper.Charlie},
			[]string{upper.Alfa, upper.Bravo},
		),
	)
}

func TestReverse(t *testing.T) {
	reversed := []string{upper.Charlie, upper.Bravo, upper.Alfa}
	strings.Reverse(reversed)
	assert.Any(t, []string{upper.Alfa, upper.Bravo, upper.Charlie}, reversed)
}

func TestSecondHalf(t *testing.T) {
	assert.Any(
		t,
		[]string{"Charlie", "Delta"},
		strings.SecondHalf(
			[]string{
				upper.Alfa,
				upper.Bravo,
				upper.Charlie,
				upper.Delta,
			},
		),
	)
}

func TestSort(t *testing.T) {
	preSorted := []string{upper.Alfa, upper.Bravo, upper.Charlie}
	strings.Sort(preSorted, true)
	assert.Strings(
		t,
		[]string{upper.Alfa, upper.Bravo, upper.Charlie},
		preSorted,
	)
	ascending := []string{upper.Bravo, upper.Alfa, upper.Charlie}
	strings.Sort(ascending, true)
	assert.Strings(
		t,
		[]string{upper.Alfa, upper.Bravo, upper.Charlie},
		ascending,
	)
	descending := []string{upper.Bravo, upper.Alfa, upper.Charlie}
	strings.Sort(descending, false)
	assert.Strings(
		t,
		[]string{upper.Charlie, upper.Bravo, upper.Alfa},
		descending,
	)
}

func TestSwap(t *testing.T) {
	actual := []string{upper.Alfa, upper.Bravo, upper.Charlie}
	strings.Swap(actual, 1, 2)
	assert.Any(t, []string{upper.Alfa, upper.Charlie, upper.Bravo}, actual)
}
