package network

import "net"

func maskedAddress(v net.HardwareAddr) bool {
	for _, b := range v[1:] {
		if b != 0 {
			return false
		}
	}

	return len(v) > 1
}
