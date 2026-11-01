package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/component"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/place_page"
	"net/http"
)

func (s *Server) placeDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	kind := component.PlaceKind(r.PathValue(constant.KindParameter))
	name := r.PathValue(constant.NameParameter)

	if kind == "" {
		s.view.RenderPage(w, name, constant.PlacesPath, place_page.Unknown())

		return
	}

	placements, e := s.service.PlacePlacements(kind, name)
	errors.PanicOnError(e)
	sightings, f := s.service.PlaceSightings(kind, name)
	errors.PanicOnError(f)
	s.view.RenderPageWithSummary(
		w,
		name,
		component.PlaceLocator(kind, name),
		place_page.Summary(kind, len(placements), len(sightings)),
		place_page.Detail(name, placements, sightings),
	)
}
