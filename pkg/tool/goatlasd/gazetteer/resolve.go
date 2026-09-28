package gazetteer

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"strings"
)

func (s *Gazetteer) Resolve(name string) (*place.Place, bool) {
	result, okay := s.place[strings.ToLower(name)]

	return result, okay
}
