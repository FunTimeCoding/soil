package mock_client

import "github.com/funtimecoding/soil/pkg/strings/join"

func key(
	collection string,
	path string,
) string {
	return join.Slash([]string{collection, path})
}
