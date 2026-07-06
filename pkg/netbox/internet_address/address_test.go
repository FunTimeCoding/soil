package internet_address

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestAddress(t *testing.T) {
	assert.NotNil(t, New(&netbox.IPAddress{Display: "10.0.0.1/24"}))
}
