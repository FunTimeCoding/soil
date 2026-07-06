package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) BoundCallsigns() []string {
	var sessions []session.Session
	errors.PanicOnError(
		s.database.Where("model_context_session != ''").Find(&sessions).Error,
	)
	var result []string

	for _, e := range sessions {
		result = append(result, e.CallsignValue())
	}

	return result
}
