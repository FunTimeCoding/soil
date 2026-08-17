package network_interface

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func New(
	device string,
	v response.NetworkInterface,
) *Interface {
	result := &Interface{
		Device:     device,
		Status:     v.Status,
		Media:      v.Media,
		MacAddress: v.MacAddress,
		Mtu:        v.Mtu,
	}

	for _, a := range v.Version4 {
		result.Addresses = append(
			result.Addresses,
			fmt.Sprintf("%s/%d", a.Value, a.Bits),
		)
	}

	for _, a := range v.Version6 {
		result.Addresses = append(
			result.Addresses,
			fmt.Sprintf("%s/%d", a.Value, a.Bits),
		)
	}

	return result
}
