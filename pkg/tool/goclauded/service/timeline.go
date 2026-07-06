package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/timeline"
	"time"
)

func (s *Service) Timeline(a argument.Timeline) ([]*timeline.Entry, error) {
	fetch := a.Limit + a.Offset
	q := event_query.New().SetLimit(fetch)

	if a.Since != "" {
		t, e := time.Parse(time.RFC3339, a.Since)

		if e == nil {
			q.SetSince(t)
		}
	}

	if a.Before != "" {
		t, e := time.Parse(time.RFC3339, a.Before)

		if e == nil {
			q.SetBefore(t)
		}
	}

	if len(a.Kinds) > 0 {
		q.SetKinds(a.Kinds)
	}

	events, e := s.store.Events(q)

	if e != nil {
		return nil, e
	}

	sessionIDs := map[string]bool{}

	for _, v := range events {
		if v.SessionIdentifier != "" {
			sessionIDs[v.SessionIdentifier] = true
		}
	}

	aliases := s.resolveAliases(sessionIDs)
	var entries []*timeline.Entry

	for _, v := range events {
		entry := timeline.FromEvent(
			v.Identifier,
			v.SessionIdentifier,
			v.Kind,
			v.Actor,
			v.Metadata,
			v.CreatedAt,
		)
		entry.Alias = aliases[v.SessionIdentifier]
		entry.Full = a.Full
		entries = append(entries, entry)
	}

	if len(a.Kinds) == 0 || hasMemoryKind(a.Kinds) {
		entries = append(entries, s.FetchVersions(a.Since, fetch)...)
	}

	return timeline.Merge(entries, a.Limit, a.Offset), nil
}
