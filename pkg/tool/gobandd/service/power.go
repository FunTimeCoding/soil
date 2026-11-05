package service

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/power_change"
)

func (s *Service) Power(
	instance string,
	operation string,
	force bool,
) (*power_change.Change, error) {
	result, e := guarded(
		s,
		instance,
		func(c face.BandClient) (*power_change.Change, error) {
			p, f := c.PowerState()

			if f != nil {
				return nil, f
			}

			state, g := Decide(operation, force, p.Available)

			if g != nil {
				return nil, g
			}

			if h := c.RequestPowerState(state); h != nil {
				return nil, h
			}

			return power_change.New(
				band.PowerStateName(p.State),
				band.PowerStateName(state),
			), nil
		},
	)

	if e != nil {
		return nil, e
	}

	s.recordIntent(instance, operation)

	return result, nil
}
