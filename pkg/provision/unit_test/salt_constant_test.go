package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/salt/constant"
	"testing"
)

func TestSaltConstant(t *testing.T) {
	assert.String(t, "SALT_HOST", constant.HostEnvironment)
}
