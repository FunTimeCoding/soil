package unit

import (
	"github.com/funtimecoding/soil/pkg/iterm/session"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/mock_client"
)

func sessionClient(identifier string) *mock_client.Client {
	result := mock_client.New()
	s := session.Stub()
	s.Identifier = identifier
	result.AddSession(s)

	return result
}
