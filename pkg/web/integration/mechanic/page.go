//go:build browser

package mechanic

import "github.com/funtimecoding/soil/pkg/strings/join"

func page(
	address string,
	path string,
) string {
	return join.Empty(address, path)
}
