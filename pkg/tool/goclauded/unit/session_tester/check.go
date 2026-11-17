package session_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (s *Session) Check() *client.CheckResponse {
	s.t.Helper()
	preview := true
	response, e := s.GeneratedClient.GetCheckWithResponse(
		s.Context,
		&client.GetCheckParams{Session: s.UUID, Preview: &preview},
	)
	assert.FatalOnError(s.t, e)

	return response.JSON200
}
