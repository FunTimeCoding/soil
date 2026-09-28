package gazetteer

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"strings"
)

func (s *Gazetteer) ResolveHardware(address string) (*place.Place, bool) {
	result, okay := s.hardware[strings.ToLower(address)]

	return result, okay
}
