package session_page

import "github.com/funtimecoding/soil/pkg/strings/join"

func Event(
	identifier string,
	part string,
) string {
	return join.Empty(Mark(identifier), "-", part)
}
