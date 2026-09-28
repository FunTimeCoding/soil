package web

import (
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
)

func placeLabel(kind string) string {
	if kind == netboxConstant.VirtualMachineAddress {
		return constant.MachineLabel
	}

	return constant.DeviceLabel
}
