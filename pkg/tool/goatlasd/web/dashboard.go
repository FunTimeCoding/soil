package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/component"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/place_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	_ *http.Request,
) {
	places, e := s.service.Places()
	errors.PanicOnError(e)
	unclaimed, f := s.service.MatchingSightings("", true)
	errors.PanicOnError(f)
	s.view.RenderPageWithSummary(
		w,
		constant.DashboardTitle,
		webConstant.RootPath,
		s.summaryItems(places, len(unclaimed)),
		place_page.Cards(places),
		html.H2(gomponents.Text(constant.UnclaimedHeading)),
		html.P(
			html.Small(
				html.A(
					html.Href(constant.SightingsPath),
					gomponents.Text(constant.AllSightingsLink),
				),
			),
		),
		component.SightingTable(unclaimed, constant.NothingUnclaimed),
	)
}
