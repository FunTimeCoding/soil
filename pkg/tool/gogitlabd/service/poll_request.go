package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/change"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_event"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
	"time"
)

func (s *Service) pollRequest(
	x context.Context,
	k *watched_request.Key,
	subscriptions []subscription.Subscription,
	due bool,
	now time.Time,
) []error {
	before, known, e := s.store.Snapshot(k)

	if e != nil {
		return []error{e}
	}

	after, d, f := s.observe(x, k)

	if not_found.Is(f) {
		return s.finish(k, subscriptions, "no longer found")
	}

	if f != nil {
		return []error{f}
	}

	actor := change.Actor(d)

	if !known {
		if g := s.baseline(x, k, after); g != nil {
			return []error{g}
		}

		return s.finishIf(k, subscriptions, after, actor)
	}

	after.Settle(before)
	after.LastNote = before.LastNote
	var events []*request_event.Event

	if after.NotesMoved(before) {
		notes, g := s.notesAfter(x, k, before.LastNote)

		if g != nil {
			return []error{g}
		}

		for _, n := range notes {
			after.LastNote = max(after.LastNote, n.Identifier)

			if n.System {
				continue
			}

			events = append(
				events,
				change.Comment(n, s.writers.Writer(n.Identifier), now),
			)
		}
	}

	for _, v := range change.Compare(before, after, actor, now) {
		events = append(events, request_event.New(v, ""))
	}

	if g := s.store.SaveSnapshot(after); g != nil {
		return []error{g}
	}

	result := s.deliver(k, subscriptions, events, due, now)

	return append(result, s.finishIf(k, subscriptions, after, actor)...)
}
