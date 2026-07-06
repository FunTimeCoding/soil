package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestReverse(t *testing.T) {
	reversed := []string{upper.Charlie, upper.Bravo, upper.Alfa}
	Reverse(reversed)
	assert.Any(t, []string{upper.Alfa, upper.Bravo, upper.Charlie}, reversed)
}
