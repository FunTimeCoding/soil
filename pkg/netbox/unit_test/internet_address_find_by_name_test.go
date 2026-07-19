package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestInternetAddressFindByName(t *testing.T) {
	address := internet_address.New(
		&netbox.IPAddress{Display: constant.FixtureAddress},
	)
	addresses := []*internet_address.Address{address}
	// Happy path
	assert.Any(
		t,
		address,
		internet_address.FindByName(addresses, constant.FixtureAddress),
	)
	// Not found
	var expected *internet_address.Address
	assert.Any(
		t,
		expected,
		internet_address.FindByName(addresses, "192.168.0.2/24"),
	)
}
