package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web/dashboard_page"
	"maragu.dev/gomponents"
)

func (s *Server) topRegion() gomponents.Node {
	records, e := s.service.TopAlerts(25, window.Whole())
	errors.PanicOnError(e)

	return dashboard_page.Top(records)
}
