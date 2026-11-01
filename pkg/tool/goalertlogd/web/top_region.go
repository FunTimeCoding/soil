package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web/dashboard_page"
	"maragu.dev/gomponents"
	"time"
)

func (s *Server) topRegion() gomponents.Node {
	now := time.Now()
	records, e := s.service.TopAlerts(25, now.Add(-7*24*time.Hour), now)
	errors.PanicOnError(e)

	return dashboard_page.Top(records)
}
