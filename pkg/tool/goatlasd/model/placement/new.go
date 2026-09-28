package placement

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/attribution"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"time"
)

func New(
	source string,
	kind string,
	scope string,
	name string,
	p *place.Place,
	seenAt time.Time,
) *Placement {
	return &Placement{
		Source:      source,
		Kind:        kind,
		Scope:       scope,
		Name:        name,
		Attribution: *attribution.New(p, seenAt),
	}
}
