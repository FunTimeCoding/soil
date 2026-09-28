package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) placeDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	kind := placeKind(r.PathValue(constant.KindParameter))
	name := r.PathValue(constant.NameParameter)

	if kind == "" {
		s.view.RenderPage(
			w,
			name,
			constant.PlacesPath,
			html.P(html.Em(gomponents.Text(constant.PlaceUnknown))),
		)

		return
	}

	placements, e := s.store.PlacePlacements(kind, name)
	errors.PanicOnError(e)
	sightings, f := s.store.PlaceSightings(kind, name)
	errors.PanicOnError(f)
	s.view.RenderPageWithSummary(
		w,
		name,
		placeLocator(kind, name),
		[]string{
			placeLabel(kind),
			Counted(
				len(placements),
				constant.PlacementWord,
				constant.PlacementsWord,
			),
			Counted(
				len(sightings),
				constant.SightingWord,
				constant.SightingsWord,
			),
		},
		html.H1(gomponents.Text(name)),
		html.H2(gomponents.Text(constant.CarriesHeading)),
		placementTable(placements, false, constant.NothingCarried),
		html.H2(gomponents.Text(constant.NetworkHeading)),
		sightingTable(sightings, constant.NothingSeen),
	)
}
