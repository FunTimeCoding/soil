package decision_page

import "github.com/funtimecoding/soil/pkg/strings/join"

func target(identifier uint) string {
	return join.Empty("#", Mark(identifier))
}
