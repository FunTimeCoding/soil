package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestTheClientClassifiesAnUnknownCallsign(t *testing.T) {
	s := base.New(t)
	_, e := s.Connector(t).PostNotify(
		t.Context(),
		client.PostNotifyJSONRequestBody{
			Callsign: "Nobody",
			Source:   "test",
			Body:     "hello",
		},
	)
	assert.True(t, not_found.Is(e))
}

func TestTheClientReadsTheStatus(t *testing.T) {
	s := base.New(t)
	_, e := s.Connector(t).GetStatus(t.Context())
	assert.FatalOnError(t, e)
}
