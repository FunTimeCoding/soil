package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/resolve_result"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func matchFrom(
	s *session.Session,
	field string,
) *resolve_result.Match {
	return resolve_result.NewMatch(
		s.Identifier,
		s.Name,
		s.AliasValue(),
		field,
	)
}
