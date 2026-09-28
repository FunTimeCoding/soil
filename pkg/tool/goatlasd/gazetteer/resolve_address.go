package gazetteer

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/place"

func (s *Gazetteer) ResolveAddress(address string) (*place.Place, bool) {
	result, okay := s.address[address]

	return result, okay
}
