package normalize

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	// Meet minimum
	assertInteger(t, 0, 0, 100, 0)
	// Below minimum
	assertInteger(t, -1, 0, 100, 0)
	// Meet maximum
	assertInteger(t, 100, 0, 100, 100)
	// Above maximum
	assertInteger(t, 101, 0, 100, 100)
	// No maximum
	assertInteger(t, 101, 0, 0, 101)
}

func assertInteger(
	t *testing.T,
	i int,
	minimum int,
	maximum int,
	expect int,
) {
	t.Helper()
	Integer(&i, minimum, maximum)
	assert.Integer(t, expect, i)
}
