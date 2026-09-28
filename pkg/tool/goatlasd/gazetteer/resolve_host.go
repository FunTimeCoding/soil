package gazetteer

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/place"

func (s *Gazetteer) ResolveHost(
	address string,
	hardwareAddress string,
) (*place.Place, bool) {
	if result, okay := s.ResolveAddress(address); okay {
		return result, true
	}

	return s.ResolveHardware(hardwareAddress)
}
