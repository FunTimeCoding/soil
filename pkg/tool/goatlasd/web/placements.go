package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) placements(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.URL.Query().Get(constant.SearchParameter)
	v, e := s.store.SearchPlacements(name)
	errors.PanicOnError(e)
	s.view.RenderPageWithSummary(
		w,
		constant.PlacementsTitle,
		constant.PlacementsPath,
		[]string{
			Counted(len(v), constant.PlacementWord, constant.PlacementsWord),
		},
		html.H1(gomponents.Text(constant.PlacementsTitle)),
		searchForm(name),
		placementTable(v, true, constant.NothingMatched),
	)
}
