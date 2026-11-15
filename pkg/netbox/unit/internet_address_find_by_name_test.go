package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/unit/netbox_tester"
	"testing"
)

func TestInternetAddressFindByName(t *testing.T) {
	address := netbox_tester.Address()
	assert.Any(
		t,
		address,
		internet_address.FindByName(
			[]*internet_address.Address{address},
			constant.FixtureAddress,
		),
	)
}

func TestInternetAddressFindByNameMissesUnknown(t *testing.T) {
	var expected *internet_address.Address
	assert.Any(
		t,
		expected,
		internet_address.FindByName(
			[]*internet_address.Address{netbox_tester.Address()},
			"192.168.0.2/24",
		),
	)
}
