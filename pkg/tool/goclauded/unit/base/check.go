package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (s *Server) Check(sessionIdentifier string) *client.CheckResponse {
	s.t.Helper()
	result, e := s.Client(s.t).GetCheck(
		s.t.Context(),
		&client.GetCheckParams{Session: sessionIdentifier},
	)
	assert.FatalOnError(s.t, e)

	return result
}
