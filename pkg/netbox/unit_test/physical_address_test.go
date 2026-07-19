package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestPhysicalAddress(t *testing.T) {
	assert.NotNil(
		t,
		physical_address.New(&netbox.MACAddress{Display: constant.PhysicalTest0}),
	)
}
