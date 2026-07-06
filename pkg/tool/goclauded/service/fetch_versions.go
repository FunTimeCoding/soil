package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/timeline"

func (s *Service) FetchVersions(
	since string,
	limit int,
) []*timeline.Entry {
	if since == "" {
		since = "2000-01-01T00:00:00Z"
	}

	versions := s.memory.VersionsSince(since, limit)
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
