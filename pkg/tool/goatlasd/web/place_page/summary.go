package place_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/component"
)

func Summary(
	kind string,
	placements int,
	sightings int,
) []string {
	return []string{
		component.PlaceLabel(kind),
		component.Counted(
			placements,
			constant.PlacementWord,
			constant.PlacementsWord,
		),
		component.Counted(
			sightings,
			constant.SightingWord,
			constant.SightingsWord,
		),
	}
}
