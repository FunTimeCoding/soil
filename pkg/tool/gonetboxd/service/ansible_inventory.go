package service

import "github.com/funtimecoding/soil/pkg/netbox/ansible_inventory"

func (s *Service) AnsibleInventory(
	field string,
) (*ansible_inventory.Inventory, error) {
	machines, e := s.client.VirtualMachines()

	if e != nil {
		return nil, e
	}

	devices, f := s.client.Devices()

	if f != nil {
		return nil, f
	}

	return ansible_inventory.New(field, machines, devices), nil
}
