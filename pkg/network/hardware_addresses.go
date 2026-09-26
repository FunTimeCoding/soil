package network

import (
	"net"
	"slices"
)

func HardwareAddresses() []string {
	interfaces, e := net.Interfaces()

	if e != nil {
		return nil
	}

	var result []string

	for _, i := range interfaces {
		if i.Flags&net.FlagLoopback != 0 || i.Flags&net.FlagUp == 0 {
			continue
		}

		if addresses, f := i.Addrs(); f != nil || len(addresses) == 0 {
			continue
		}

		v := i.HardwareAddr.String()

		if v == "" || maskedAddress(i.HardwareAddr) ||
			slices.Contains(result, v) {
			continue
		}

		result = append(result, v)
	}

	return result
}
