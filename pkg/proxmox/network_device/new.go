package network_device

import (
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"net"
	"strconv"
	"strings"
)

func New(
	name string,
	value string,
) *Device {
	result := &Device{Name: name}

	for _, p := range strings.Split(value, stringsConstant.Comma) {
		k, v, okay := strings.Cut(p, stringsConstant.Equals)

		if !okay {
			continue
		}

		switch k {
		case constant.BridgeKey:
			result.Bridge = v
		case constant.Name:
			result.Interface = v
		case constant.HardwareAddressKey:
			result.HardwareAddress = v
		case constant.TagKey:
			vlan, e := strconv.Atoi(v)

			if e == nil {
				result.Vlan = vlan
			}
		default:
			if _, e := net.ParseMAC(v); e == nil {
				result.Model = k
				result.HardwareAddress = v
			}
		}
	}

	return result
}
