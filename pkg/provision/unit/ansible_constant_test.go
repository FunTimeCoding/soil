package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"testing"
)

func TestAnsibleConstant(t *testing.T) {
	assert.String(t, "ANSIBLE_INVENTORY", constant.AnsibleInventoryEnvironment)
}
