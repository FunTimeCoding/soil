package decision_page

import "github.com/funtimecoding/soil/pkg/strings/join"

func Event(
	identifier uint,
	part string,
) string {
	return join.Empty(Mark(identifier), "-", part)
}
