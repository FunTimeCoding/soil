package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) raidDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, e := strconv.Atoi(r.PathValue("id"))
	errors.PanicOnError(e)
	fights := s.store.RaidFights(identifier)
	players := s.store.RaidPlayerStats(identifier)
	s.view.RenderPage(
		w,
		"Raid",
		constant.RaidsPath,
		html.H1(gomponents.Textf("Raid - %d fights", len(fights))),
		raidDetailTable(players),
	)
}
