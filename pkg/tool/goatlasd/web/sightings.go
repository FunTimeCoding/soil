package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
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
	v, e := s.store.MatchingSightings("", unclaimed)
	errors.PanicOnError(e)
	s.view.RenderPageWithSummary(
		w,
		constant.SightingsTitle,
		constant.SightingsPath,
		[]string{
			Counted(len(v), constant.SightingWord, constant.SightingsWord),
		},
		html.H1(gomponents.Text(constant.SightingsTitle)),
		html.P(html.Small(sightingToggle(unclaimed))),
		sightingTable(v, constant.NothingUnclaimed),
	)
}
