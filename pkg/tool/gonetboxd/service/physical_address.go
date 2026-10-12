package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net"
)

func physicalAddress(s string) (net.HardwareAddr, error) {
	result, e := net.ParseMAC(s)

	if e != nil {
		return nil, validation.New("physical address %s: %v", s, e)
	}

	return result, nil
}
