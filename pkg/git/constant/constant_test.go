package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "v", VersionPrefix)
	assert.String(t, "HEAD", HeadReference)
	assert.Integer(t, 7, HashLength)
}
