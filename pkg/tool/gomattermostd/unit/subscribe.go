package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"testing"
)

func subscribe(
	t *testing.T,
	s *service.Service,
	m *mock_client.Client,
	root string,
	callsign string,
	alias string,
) {
	t.Helper()
	p, e := m.FindPost(root)
	assert.FatalOnError(t, e)
	_, _, f := s.Subscribe(p, callsign, alias)
	assert.FatalOnError(t, f)
}
