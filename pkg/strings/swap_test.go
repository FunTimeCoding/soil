package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestSwapInList(t *testing.T) {
	actual := []string{upper.Alfa, upper.Bravo, upper.Charlie}
	Swap(actual, 1, 2)
	assert.Any(t, []string{upper.Alfa, upper.Charlie, upper.Bravo}, actual)
}
