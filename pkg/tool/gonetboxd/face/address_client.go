package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

type AddressClient interface {
	CreateAddress(
		x context.Context,
		device string,
		interfaceName string,
		address string,
		status string,
	) (client.Address, error)
	CreateAddressRange(
		x context.Context,
		start string,
		end string,
		status string,
		description string,
	) (client.AddressRange, error)
	CreatePrefix(
		x context.Context,
		prefix string,
		site string,
		description string,
	) (client.Prefix, error)
	DeleteAddress(
		x context.Context,
		identifier int32,
	) error
	ListAddressRanges(
		x context.Context,
	) ([]*client.AddressRange, error)
	ListAddresses(
		x context.Context,
		name string,
	) ([]*client.Address, error)
	ListPhysicalAddresses(
		x context.Context,
	) ([]client.PhysicalAddressOwner, error)
	ListPrefixes(
		x context.Context,
	) ([]*client.Prefix, error)
	SetInterfacePhysicalAddress(
		x context.Context,
		device string,
		name string,
		address string,
	) (client.PhysicalAddress, error)
}
