package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (s *Server) deliver(
	name string,
	to string,
	body string,
	d client.Delivery,
) {
	s.t.Helper()
	request := client.PostSendJSONRequestBody{
		Callsign: name,
		Body:     body,
		Delivery: &d,
	}

	if to != "" {
		request.To = &[]string{to}
	}

	_, e := s.Client(s.t).PostSend(s.t.Context(), request)
	assert.FatalOnError(s.t, e)
}
