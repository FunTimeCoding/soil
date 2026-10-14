package model_context_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (s *Session) Label(key string) string {
	s.T.Helper()
	response, e := s.RestClient.GetSessionLabelWithResponse(s.Context, s.UUID)
	assert.FatalOnError(s.T, e)

	for _, l := range response.JSON200.Labels {
		if l.Key == key {
			return l.Value
		}
	}

	return ""
}
