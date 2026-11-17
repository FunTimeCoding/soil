package session_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (s *Session) Label(key string) string {
	s.t.Helper()
	response, e := s.GeneratedClient.GetSessionLabelWithResponse(
		s.Context,
		s.UUID,
	)
	assert.FatalOnError(s.t, e)

	for _, l := range response.JSON200.Labels {
		if l.Key == key {
			return l.Value
		}
	}

	return ""
}
