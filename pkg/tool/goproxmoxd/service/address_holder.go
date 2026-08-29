package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/network_device"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"strings"
)

func (s *Service) addressHolder(
	c face.ProxmoxClient,
	address string,
	exclude int,
) (*int, error) {
	nodes, e := c.Nodes()

	if e != nil {
		return nil, e
	}

	for _, status := range nodes {
		node, f := c.Node(status.Node)

		if f != nil {
			return nil, f
		}

		machines, g := c.Machines(node)

		if g != nil {
			return nil, g
		}

		for _, v := range machines {
			current := int(v.VMID)

			if current == exclude {
				continue
			}

			detail, h := c.Machine(node, current)

			if h != nil {
				return nil, h
			}

			if detail.VirtualMachineConfig == nil {
				continue
			}

			for _, d := range network_device.NewSlice(
				detail.VirtualMachineConfig.Nets,
			) {
				if strings.EqualFold(d.HardwareAddress, address) {
					return &current, nil
				}
			}
		}
	}

	return nil, nil
}
