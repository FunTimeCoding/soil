package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "v", constant.VersionPrefix)
	assert.String(t, "HEAD", constant.HeadReference)
	assert.Integer(t, 7, constant.HashLength)
}
