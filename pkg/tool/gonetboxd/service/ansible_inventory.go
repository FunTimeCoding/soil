package service

import "github.com/funtimecoding/soil/pkg/netbox/ansible_inventory"

func (s *Service) AnsibleInventory(
	field string,
	site string,
) (*ansible_inventory.Inventory, error) {
	return guarded(
		s.guard,
		func() (*ansible_inventory.Inventory, error) {
			machines, e := s.client.VirtualMachines()

			if e != nil {
				return nil, e
			}

			devices, f := s.client.Devices()

			if f != nil {
				return nil, f
			}

			if s.powerIntent {
				machines = activeMachines(machines)
				devices = activeDevices(devices)
			}

			tunnel, g := s.tunnelAddresses(site, machines, devices)

			if g != nil {
				return nil, g
			}

			return ansible_inventory.New(field, tunnel, machines, devices), nil
		},
	)
}
