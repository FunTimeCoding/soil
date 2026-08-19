package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	_ *http.Request,
) {
	authority, e := s.store.Authorities()
	errors.PanicOnError(e)
	pending, f := s.service.Pending()
	errors.PanicOnError(f)
	horizon := store.NewFilter()
	horizon.Before = new(time.Now().AddDate(0, 0, constant.ExpiryHorizonDay))
	horizon.Revoked = new(false)
	expiring, g := s.store.Certificates(horizon)
	errors.PanicOnError(g)
	s.view.RenderPage(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		html.H1(gomponents.Text(constant.DashboardTitle)),
		html.Div(
			html.Class("summary-cards"),
			card(constant.AuthoritiesTitle, len(authority)),
			card("Pending publication", len(pending)),
			card("Expiring within a month", len(expiring)),
		),
		html.H2(gomponents.Text(constant.AuthoritiesTitle)),
		authoritiesTable(authority),
		publishForm(pending),
	)
}
