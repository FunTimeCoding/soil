package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/ansible_inventory"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) AnsibleInventory(
	field string,
	site string,
) (*ansible_inventory.Inventory, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*ansible_inventory.Inventory, error) {
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
