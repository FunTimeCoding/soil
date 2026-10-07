package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/match"
)

func matchFrom(
	s *session.Session,
	field string,
) *match.Match {
	return match.New(s.Identifier, s.Name, s.AliasValue(), field)
}
