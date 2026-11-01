package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/filter"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/web/authority_page"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/web/dashboard_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	_ *http.Request,
) {
	authority, e := s.service.Authorities()
	errors.PanicOnError(e)
	pending, f := s.service.Pending()
	errors.PanicOnError(f)
	horizon := filter.New()
	horizon.Before = new(time.Now().AddDate(0, 0, constant.ExpiryHorizonDay))
	horizon.Revoked = new(false)
	expiring, g := s.service.Certificates(horizon)
	errors.PanicOnError(g)
	s.view.RenderPageWithSummary(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		dashboard_page.Summary(len(authority), len(pending), len(expiring)),
		html.H1(gomponents.Text(constant.DashboardTitle)),
		html.H2(gomponents.Text(constant.AuthoritiesTitle)),
		authority_page.Table(authority),
		dashboard_page.Publish(pending),
	)
}
