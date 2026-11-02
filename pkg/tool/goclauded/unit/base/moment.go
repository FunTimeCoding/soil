package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Server) Moment(
	sessionIdentifier string,
	name string,
	line string,
) {
	s.t.Helper()
	assert.FatalOnError(
		s.t,
		s.Store.Store.LogEvent(
			sessionIdentifier,
			constant.Moment,
			name,
			map[string]string{"line": line},
		),
	)
}
