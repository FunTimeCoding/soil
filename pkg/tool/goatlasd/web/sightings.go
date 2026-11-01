package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/component"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/sighting_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) sightings(
	w http.ResponseWriter,
	r *http.Request,
) {
	unclaimed := r.URL.Query().Get(constant.UnclaimedParameter) ==
		constant.UnclaimedValue
	v, e := s.service.MatchingSightings("", unclaimed)
	errors.PanicOnError(e)
	s.view.RenderPageWithSummary(
		w,
		constant.SightingsTitle,
		constant.SightingsPath,
		[]string{
			component.Counted(
				len(v),
				constant.SightingWord,
				constant.SightingsWord,
			),
		},
		html.H1(gomponents.Text(constant.SightingsTitle)),
		html.P(html.Small(sighting_page.Toggle(unclaimed))),
		component.SightingTable(v, constant.NothingUnclaimed),
	)
}
