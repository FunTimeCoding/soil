package session_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (s *Session) CheckLive() *client.CheckResponse {
	s.t.Helper()
	response, e := s.GeneratedClient.GetCheckWithResponse(
		s.Context,
		&client.GetCheckParams{Session: s.UUID},
	)
	assert.FatalOnError(s.t, e)

	return response.JSON200
}
