package fetch

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFetch(t *testing.T) {
	assert.NotNil(t, Command())
}
