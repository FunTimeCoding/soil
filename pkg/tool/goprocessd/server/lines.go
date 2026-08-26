package server

import "github.com/funtimecoding/soil/pkg/strings/join"

func lines(history []string) string {
	if len(history) == 0 {
		return "ok"
	}

	return join.NewLine(history)
}
