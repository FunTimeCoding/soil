package placement

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/attribution"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"time"
)

func NewPackage(
	source string,
	kind string,
	scope string,
	name string,
	packageName string,
	version string,
	p *place.Place,
	seenAt time.Time,
) *Placement {
	return &Placement{
		Source:      source,
		Kind:        kind,
		Scope:       scope,
		Name:        name,
		Package:     packageName,
		Version:     version,
		Attribution: *attribution.New(p, seenAt),
	}
}
