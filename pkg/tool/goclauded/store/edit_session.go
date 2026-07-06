package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument/edit_session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) EditSession(
	identifier string,
	a *edit_session.Session,
) error {
	updates := map[string]any{}

	if a.Alias != nil {
		if *a.Alias == "" {
			updates[constant.Alias] = nil
		} else {
			updates[constant.Alias] = *a.Alias
		}
	}

	if a.Description != nil {
		updates[constant.Description] = *a.Description
	}

	if a.Slug != nil {
		updates[constant.Slug] = *a.Slug
	}

	if a.Topic != nil {
		updates[constant.Topic] = *a.Topic
	}

	if a.Files != nil {
		updates[constant.Files] = *a.Files
	}

	if len(updates) == 0 {
		return nil
	}

	updates["last_seen"] = s.clock()

	return s.database.Model(session.Stub()).Where(
		"identifier = ?",
		identifier,
	).Updates(updates).Error
}
