package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) AliasedSessions() []session.Session {
	var result []session.Session
	errors.PanicOnError(
		s.database.
			Where(fmt.Sprintf("%s != ''", constant.Alias)).
			Order(constant.Alias).
			Find(&result).Error,
	)

	return result
}
