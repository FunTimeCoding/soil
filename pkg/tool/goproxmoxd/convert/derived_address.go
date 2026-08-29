package convert

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"

func DerivedAddress(
	instance string,
	identifier int,
	address string,
	holder *int,
) *server.DerivedAddress {
	return &server.DerivedAddress{
		Instance:        instance,
		Identifier:      identifier,
		HardwareAddress: address,
		InUseBy:         holder,
	}
}
