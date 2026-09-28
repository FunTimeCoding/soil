package attribution

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/place"
	"time"
)

func New(
	p *place.Place,
	seenAt time.Time,
) *Attribution {
	result := &Attribution{SeenAt: seenAt}

	if p == nil {
		return result
	}

	result.PlaceKind = p.Kind
	result.PlaceIdentifier = p.Identifier
	result.PlaceName = p.Name

	return result
}
