package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/summary"
)

func (s *Store) UpdateSummaryBody(
	sessionIdentifier string,
	body string,
) {
	errors.PanicOnError(
		s.database.Model(summary.Stub()).
			Where("session_identifier = ?", sessionIdentifier).
			Update(constant.Body, body).Error,
	)
}
