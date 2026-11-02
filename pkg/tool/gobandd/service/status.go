package service

import (
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/status"
)

func (s *Service) Status(instance string) (*status.Status, error) {
	return guarded(
		s,
		instance,
		func(c face.BandClient) (*status.Status, error) {
			result := status.New(instance)

			for _, read := range []func(face.BandClient, *status.Status) error{
				readHost,
				readSetup,
				readScreen,
				readRedirection,
				readScreenService,
				readConsent,
				readPower,
			} {
				if e := read(c, result); e != nil {
					return nil, e
				}
			}

			return result, nil
		},
	)
}
