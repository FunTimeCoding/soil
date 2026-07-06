package separator

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "#", Hash)
	assert.String(t, "//", DoubleSlash)
}
