package network_device

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"net"
)

func Derive(
	instance int,
	identifier int,
) (string, error) {
	if instance < 0 || instance > constant.MaximumInstance {
		return "", validation.New("instance index out of range: %d", instance)
	}

	if identifier < 0 || identifier > constant.MaximumMachine {
		return "", validation.New(
			"machine identifier out of range: %d",
			identifier,
		)
	}

	return net.HardwareAddr{
		constant.LocalPrefix,
		byte(instance),
		byte(identifier >> 24),
		byte(identifier >> 16),
		byte(identifier >> 8),
		byte(identifier),
	}.String(), nil
}
