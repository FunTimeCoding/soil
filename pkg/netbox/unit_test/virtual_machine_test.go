package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestMachine(t *testing.T) {
	assert.NotNil(
		t,
		virtual_machine.New(&netbox.VirtualMachineWithConfigContext{}),
	)
}
