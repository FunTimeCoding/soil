package mock_client

import "github.com/funtimecoding/soil/pkg/strings/join"

func key(parts ...string) string {
	return join.Slash(parts)
}
