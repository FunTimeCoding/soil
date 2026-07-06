package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) RefreshSession(
	identifier string,
	state *tracker.State,
) {
	updates := map[string]any{}

	if state.Slug != "" {
		updates[constant.Slug] = state.Slug
	}

	if state.WorkDirectory != "" {
		updates["work_directory"] = state.WorkDirectory
	}

	if state.Branch != "" {
		updates["branch"] = state.Branch
	}

	if state.FirstTimestamp != "" {
		updates["session_timestamp"] = state.FirstTimestamp

		if t, e := time.Parse(
			time.RFC3339Nano,
			state.FirstTimestamp,
		); e == nil {
			updates["started_at"] = t
		}
	}

	if state.LastTimestamp != "" {
		if t, e := time.Parse(
			time.RFC3339Nano,
			state.LastTimestamp,
		); e == nil {
			updates["last_active_at"] = t
		}
	}

	if state.Lines > 0 {
		updates["lines"] = state.Lines
	}

	if state.UserMessageCount > 0 {
		updates["turn_count"] = state.UserMessageCount
	}

	if state.FirstMessage != "" {
		updates["first_message"] = state.FirstMessage
	}

	if len(updates) > 0 {
		s.store.UpdateFields(identifier, updates)
	}
}
