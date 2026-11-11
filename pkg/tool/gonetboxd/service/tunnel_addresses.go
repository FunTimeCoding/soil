package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (s *Service) tunnelAddresses(
	site string,
	machines []*virtual_machine.Machine,
	devices []*device.Device,
) (map[string]string, error) {
	result := map[string]string{}

	if site == "" {
		return result, nil
	}

	terminations, e := s.client.TunnelTerminations()

	if e != nil {
		return nil, e
	}

	ends := map[string]bool{}

	for _, t := range terminations {
		ends[terminationKey(t.TerminationType, t.TerminationIdentifier)] = true
	}

	for _, d := range devices {
		if d.Site == site {
			continue
		}

		addresses, f := s.client.DeviceAddresses(d.Name)

		if f != nil {
			return nil, f
		}

		if a := tunnelAddress(ends, addresses); a != "" {
			result[d.Name] = a
		}
	}

	for _, m := range machines {
		if m.Site == site {
			continue
		}

		addresses, f := s.client.VirtualMachineAddresses(m.Name)

		if f != nil {
			return nil, f
		}

		if a := tunnelAddress(ends, addresses); a != "" {
			result[m.Name] = a
		}
	}

	return result, nil
}
