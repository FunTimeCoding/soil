package netbox_tester

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func Alfa() *virtual_machine.Machine {
	result := virtual_machine.Stub()
	result.Identifier = 2
	result.Name = "alfa"

	return result
}
