package unit

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/netbox-community/go-netbox/v4"
)

func newFixtureAddress() *internet_address.Address {
	return internet_address.New(
		&netbox.IPAddress{Display: constant.FixtureAddress},
	)
}
