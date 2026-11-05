package unit

import (
	"github.com/funtimecoding/soil/pkg/iterm/mock_client"
	"github.com/funtimecoding/soil/pkg/iterm/session"
)

func sessionClient(identifier string) *mock_client.Client {
	result := mock_client.New()
	s := session.Stub()
	s.Identifier = identifier
	result.AddSession(s)

	return result
}
