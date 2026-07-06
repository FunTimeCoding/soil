package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestSendDirectMessage(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	b.CheckLive()
	a.MustCallTool(
		constant.Send,
		map[string]any{
			constant.To:   b.Name(),
			constant.Body: "heads up: running lint",
		},
	)
	r := b.CheckLive()
	messages := clientEntriesByKind(r.Entries, constant.QueueMessage)
	assert.Count(t, 1, messages)
	assert.StringContains(t, a.Name(), messages[0].Body)
	assert.StringContains(t, "heads up: running lint", messages[0].Body)
}

func TestSendBroadcast(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "listener")
	b.CheckLive()
	a.MustCallTool(
		constant.Send,
		map[string]any{
			constant.Body: "deploying service",
		},
	)
	r := b.CheckLive()
	messages := clientEntriesByKind(r.Entries, constant.QueueMessage)
	assert.Count(t, 1, messages)
	assert.StringContains(t, "deploying service", messages[0].Body)
}
