package fixture

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFixture(t *testing.T) {
	assert.Suffix(
		t,
		"soil/fixture/example.txt",
		Path("example.txt"),
	)
}
