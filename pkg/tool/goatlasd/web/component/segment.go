package component

import (
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
)

func segment(kind string) string {
	if kind == netboxConstant.VirtualMachineAddress {
		return constant.MachineSegment
	}

	return constant.DeviceSegment
}
