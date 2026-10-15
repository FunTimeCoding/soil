package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/machine_change"
)

func (s *Service) UpdateVirtualMachine(
	name string,
	c *machine_change.Change,
) (*virtual_machine.Machine, error) {
	if c.PrimaryAddress != nil {
		if _, e := s.client.SetVirtualMachinePrimaryAddress(
			name,
			*c.PrimaryAddress,
		); e != nil {
			return nil, e
		}
	}

	if c.Platform != nil {
		if _, e := s.client.SetVirtualMachinePlatform(
			name,
			*c.Platform,
		); e != nil {
			return nil, e
		}
	}

	if c.Role != nil {
		if _, e := s.client.SetVirtualMachineRole(name, *c.Role); e != nil {
			return nil, e
		}
	}

	if c.Tenant != nil {
		if _, e := s.client.SetVirtualMachineTenant(name, *c.Tenant); e != nil {
			return nil, e
		}
	}

	if c.Cores != nil {
		if _, e := s.client.SetVirtualMachineCores(name, *c.Cores); e != nil {
			return nil, e
		}
	}

	if c.Memory != nil {
		if _, e := s.client.SetVirtualMachineMemory(name, *c.Memory); e != nil {
			return nil, e
		}
	}

	if c.Status != nil {
		if _, e := s.client.SetVirtualMachineStatus(name, *c.Status); e != nil {
			return nil, e
		}
	}

	if c.CustomFields != nil {
		if _, e := s.client.SetVirtualMachineCustomFields(
			name,
			c.CustomFields,
		); e != nil {
			return nil, e
		}
	}

	if c.Name != nil {
		renamed, e := s.client.RenameVirtualMachine(name, *c.Name)

		if e != nil {
			return nil, e
		}

		name = renamed.Name
	}

	return s.client.VirtualMachineByName(name)
}
