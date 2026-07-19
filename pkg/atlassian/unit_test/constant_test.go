package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "ATLASSIAN_HOST", constant.HostEnvironment)
}
