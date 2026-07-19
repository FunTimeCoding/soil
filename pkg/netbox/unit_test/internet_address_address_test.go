package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestInternetAddressAddress(t *testing.T) {
	assert.NotNil(
		t,
		internet_address.New(&netbox.IPAddress{Display: "10.0.0.1/24"}),
	)
}
