package internet_address

import (
	"net"
	"strings"
)

func Find(addresses []*Address, address string) *Address {
	bare, _, _ := strings.Cut(address, "/")
	target := net.ParseIP(bare)

	if target == nil {
		return nil
	}

	for _, a := range addresses {
		if target.Equal(a.Address) {
			return a
		}
	}

	return nil
}
