package physical_address

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestAddress(t *testing.T) {
	assert.NotNil(t, New(&netbox.MACAddress{Display: constant.PhysicalTest0}))
}
