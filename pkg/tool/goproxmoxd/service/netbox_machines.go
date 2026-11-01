package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (s *Service) netboxMachines(
	x context.Context,
) (map[string][]*client.VirtualMachine, error) {
	var listed []*client.VirtualMachine
	e := s.netboxGuard.Attempt(
		x,
		func(y context.Context) error {
			var f error
			listed, f = s.netbox.ListVirtualMachines(y)

			return f
		},
	)

	if e != nil {
		return nil, e
	}

	result := map[string][]*client.VirtualMachine{}

	for _, v := range listed {
		result[v.Name] = append(result[v.Name], v)
	}

	return result, nil
}
