package place_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/component"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Detail(
	name string,
	placements []*placement.Placement,
	sightings []*sighting.Sighting,
) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H1(gomponents.Text(name)),
			html.H2(gomponents.Text(constant.CarriesHeading)),
			component.PlacementTable(placements, false, constant.NothingCarried),
			html.H2(gomponents.Text(constant.NetworkHeading)),
			component.SightingTable(sightings, constant.NothingSeen),
		},
	)
}
