package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/raid_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) raidDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, e := strconv.Atoi(r.PathValue(constant.IdentifierParameter))
	errors.PanicOnError(e)
	fights := s.service.RaidFights(identifier)
	players := s.service.RaidPlayerStats(identifier)
	s.view.RenderPage(
		w,
		"Raid",
		constant.RaidsPath,
		html.H1(gomponents.Textf("Raid - %d fights", len(fights))),
		raid_page.Detail(players),
	)
}
