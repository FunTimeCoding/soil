package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestSort(t *testing.T) {
	preSorted := []string{upper.Alfa, upper.Bravo, upper.Charlie}
	Sort(preSorted, true)
	assert.Strings(
		t,
		[]string{upper.Alfa, upper.Bravo, upper.Charlie},
		preSorted,
	)
	ascending := []string{upper.Bravo, upper.Alfa, upper.Charlie}
	Sort(ascending, true)
	assert.Strings(
		t,
		[]string{upper.Alfa, upper.Bravo, upper.Charlie},
		ascending,
	)
	descending := []string{upper.Bravo, upper.Alfa, upper.Charlie}
	Sort(descending, false)
	assert.Strings(
		t,
		[]string{upper.Charlie, upper.Bravo, upper.Alfa},
		descending,
	)
}
