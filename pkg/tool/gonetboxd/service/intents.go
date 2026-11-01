package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/intent"
)

func (s *Service) Intents() ([]*intent.Intent, error) {
	devices, e := s.client.Devices()

	if e != nil {
		return nil, e
	}

	machines, f := s.client.VirtualMachines()

	if f != nil {
		return nil, f
	}

	result := make([]*intent.Intent, 0, len(devices)+len(machines))

	for _, d := range devices {
		result = append(
			result,
			intent.New(constant.DeviceKind, d.Name, d.Status),
		)
	}

	for _, m := range machines {
		result = append(
			result,
			intent.New(constant.MachineKind, m.Name, m.Status),
		)
	}

	return result, nil
}
