package gazetteer

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"strings"
)

func (s *Gazetteer) add(
	p *place.Place,
	primary *string,
) {
	s.place[strings.ToLower(p.Name)] = p

	if primary == nil || *primary == "" {
		return
	}

	s.address[strings.SplitN(*primary, "/", 2)[0]] = p
}
