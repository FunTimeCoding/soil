package session_tester

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Session) SendToolNotation() string {
	s.t.Helper()

	for _, tool := range s.ListTools() {
		if tool.Name == constant.Send {
			return string(notation.Marshal(tool))
		}
	}

	s.t.Fatalf("no %s tool listed", constant.Send)

	return ""
}
