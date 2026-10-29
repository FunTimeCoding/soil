package service

import "github.com/funtimecoding/soil/pkg/habitica/gear"

func (s *Service) UserGear() (*gear.Gear, error) {
	return guarded(s.guard, s.client.UserGear)
}
