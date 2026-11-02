package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (s *Server) Announce(
	name string,
	topic string,
) {
	s.t.Helper()
	assert.FatalOnError(
		s.t,
		s.Client(s.t).PostAnnounce(
			s.t.Context(),
			client.PostAnnounceJSONRequestBody{Callsign: name, Topic: topic},
		),
	)
}
