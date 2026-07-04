package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}
