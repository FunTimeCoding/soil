package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func activeMachines(v []*virtual_machine.Machine) []*virtual_machine.Machine {
	var result []*virtual_machine.Machine

	for _, m := range v {
		if m.Status == constant.StatusActive {
			result = append(result, m)
		}
	}

	return result
}
