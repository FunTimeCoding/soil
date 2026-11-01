package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store/result"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/component"
)

func (*Server) summaryItems(
	places []*result.Place,
	unclaimed int,
) []string {
	placements := 0

	for _, p := range places {
		placements += p.Count
	}

	return []string{
		component.Counted(len(places), constant.PlaceWord, constant.PlacesWord),
		component.Counted(
			placements,
			constant.PlacementWord,
			constant.PlacementsWord,
		),
		fmt.Sprintf(constant.UnclaimedSummary, unclaimed),
	}
}
