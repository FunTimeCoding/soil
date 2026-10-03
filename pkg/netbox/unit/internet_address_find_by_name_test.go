package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"testing"
)

func TestInternetAddressFindByName(t *testing.T) {
	address := newFixtureAddress()
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
			[]*internet_address.Address{newFixtureAddress()},
			"192.168.0.2/24",
		),
	)
}
