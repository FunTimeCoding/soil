package teams_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/teams"
	"github.com/funtimecoding/soil/pkg/microsoft/graph_server"
	"github.com/funtimecoding/soil/pkg/microsoft/graph_tester"
	"github.com/funtimecoding/soil/pkg/microsoft/mock_token_store"
	"testing"
)

func New(
	t *testing.T,
	s *graph_server.Server,
) *teams.Client {
	t.Helper()
	result := teams.New(
		graph_tester.New(t).Client(
			s,
			mock_token_store.New(t, graph_tester.Valid()),
		),
	)
	t.Cleanup(result.Close)

	return result
}
