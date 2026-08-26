package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Server) handleStatus() string {
	var result []string

	for _, e := range s.Statuses() {
		prefix := " "

		if e.Running {
			prefix = "*"
		}

		result = append(result, fmt.Sprintf("%s%s", prefix, e.Name))
	}

	return join.NewLine(result)
}
