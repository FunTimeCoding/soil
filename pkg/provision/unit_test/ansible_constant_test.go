package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/ansible/constant"
	"testing"
)

func TestAnsibleConstant(t *testing.T) {
	assert.String(t, "ANSIBLE_INVENTORY", constant.InventoryEnvironment)
}
