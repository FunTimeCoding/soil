package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) ReleaseByCallsign(c string) error {
	return s.database.Model(session.Stub()).Where(
		"callsign = ?",
		c,
	).Update(constant.Callsign, nil).Error
}
