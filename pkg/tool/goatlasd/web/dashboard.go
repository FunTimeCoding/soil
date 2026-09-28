package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	_ *http.Request,
) {
	places, e := s.store.Places()
	errors.PanicOnError(e)
	unclaimed, f := s.store.MatchingSightings("", true)
	errors.PanicOnError(f)
	s.view.RenderPageWithSummary(
		w,
		constant.DashboardTitle,
		webConstant.RootPath,
		s.summaryItems(places, len(unclaimed)),
		html.Div(
			html.Class(constant.CardGridClass),
			gomponents.Map(places, placeCard),
		),
		html.H2(gomponents.Text(constant.UnclaimedHeading)),
		html.P(
			html.Small(
				html.A(
					html.Href(constant.SightingsPath),
					gomponents.Text(constant.AllSightingsLink),
				),
			),
		),
		sightingTable(unclaimed, constant.NothingUnclaimed),
	)
}
