package web

import (
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
)

func placeKind(segment string) string {
	switch segment {
	case constant.DeviceSegment:
		return netboxConstant.DeviceAddress
	case constant.MachineSegment:
		return netboxConstant.VirtualMachineAddress
	}

	return ""
}
