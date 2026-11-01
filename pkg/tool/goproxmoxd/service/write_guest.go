package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (s *Service) writeGuest(
	x context.Context,
	name string,
	status string,
) error {
	return s.netboxGuard.Attempt(
		x,
		func(y context.Context) error {
			body := client.UpdateVirtualMachineRequest{Status: &status}
			_, f := s.netbox.UpdateVirtualMachine(y, name, body)

			return f
		},
	)
}
