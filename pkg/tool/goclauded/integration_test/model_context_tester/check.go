package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (s *Session) Check() *client.CheckResponse {
	s.T.Helper()
	preview := true
	response, e := s.RestClient.GetCheckWithResponse(
		s.Context,
		&client.GetCheckParams{
			Session: s.UUID,
			Preview: &preview,
		},
	)
	assert.FatalOnError(s.T, e)

	return response.JSON200
}
