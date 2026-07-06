package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Server) handleStatus() string {
	var lines []string

	for _, p := range s.processes {
		prefix := " "

		if p.Running() {
			prefix = "*"
		}

		lines = append(lines, fmt.Sprintf("%s%s", prefix, p.Name))
	}

	return join.NewLine(lines)
}
