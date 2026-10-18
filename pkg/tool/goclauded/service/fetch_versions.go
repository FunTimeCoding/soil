package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/timeline"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/version_entry"
	"time"
)

func (s *Service) FetchVersions(
	x context.Context,
	since string,
	limit int,
) []*timeline.Entry {
	if since == "" {
		since = "2000-01-01T00:00:00Z"
	}

	var versions []version_entry.Entry

	if e := s.memoryGuard.Do(
		x,
		func(y context.Context) error {
			var f error
			versions, f = s.memory.VersionsSince(y, since, limit)

			return f
		},
	); e != nil {
		return []*timeline.Entry{
			timeline.FromUnreachable(e, s.clock().UTC().Format(time.RFC3339)),
		}
	}
	var entries []*timeline.Entry

	for _, v := range versions {
		entries = append(
			entries,
			timeline.FromVersion(
				v.Name,
				v.Description,
				v.ChangeType,
				v.Source,
				v.ChangedAt,
			),
		)
	}

	return entries
}
