package fixture

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFixture(t *testing.T) {
	assert.Suffix(
		t,
		"go-library/fixture/example.txt",
		Path("example.txt"),
	)
}
