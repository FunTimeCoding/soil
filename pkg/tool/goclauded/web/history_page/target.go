package history_page

import "github.com/funtimecoding/soil/pkg/strings/join"

func target(identifier uint) string {
	return join.Empty("#", mark(identifier))
}
