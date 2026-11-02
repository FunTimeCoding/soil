package service

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/face"
)

func (s *Service) Power(
	instance string,
	operation string,
	force bool,
) (string, error) {
	result, e := guarded(
		s,
		instance,
		func(c face.BandClient) (string, error) {
			p, f := c.PowerState()

			if f != nil {
				return "", f
			}

			state, g := Decide(operation, force, p.Available)

			if g != nil {
				return "", g
			}

			if h := c.RequestPowerState(state); h != nil {
				return "", h
			}

			return band.PowerStateName(state), nil
		},
	)

	if e != nil {
		return "", e
	}

	s.recordIntent(instance, operation)

	return result, nil
}
