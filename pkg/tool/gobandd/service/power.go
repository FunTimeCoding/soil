package service

import "github.com/funtimecoding/soil/pkg/band"

func (s *Service) Power(
	instance string,
	operation string,
	force bool,
) (string, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return "", e
	}

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

	s.recordIntent(instance, operation)

	return band.PowerStateName(state), nil
}
