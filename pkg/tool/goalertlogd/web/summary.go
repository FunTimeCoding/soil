package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web/dashboard_page"
)

func (s *Server) summary() []string {
	status, e := s.service.Status()
	errors.PanicOnError(e)

	return dashboard_page.Summary(status, s.worker.LastPoll())
}
