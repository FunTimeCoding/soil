package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_event"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
	"github.com/funtimecoding/soil/pkg/watch/digest"
	"time"
)

func (s *Service) deliver(
	k *watched_request.Key,
	subscriptions []subscription.Subscription,
	events []*request_event.Event,
	due bool,
	now time.Time,
) []error {
	root := k.String()
	identifiers := make([]uint, 0, len(subscriptions))

	for _, v := range subscriptions {
		identifiers = append(identifiers, v.Identifier)
	}

	s.pending.Retain(root, identifiers)
	var result []error

	for _, v := range subscriptions {
		visible := request_event.Visible(events, v.Callsign)

		if len(visible) == 0 && !due {
			continue
		}

		combined := append(s.pending.Take(v.Identifier), visible...)

		if len(combined) == 0 {
			continue
		}

		e := s.digests.Notify(
			v.Callsign,
			digest.Build(v.Label(), combined),
			v.DeliveryMode(),
		)
		refused := s.note(v.Callsign, e)

		if e == nil {
			s.pending.Done(v.Identifier)

			continue
		}

		s.pending.Keep(v.Identifier, root, combined, now)

		if refused != nil {
			result = append(result, refused)
		}
	}

	return result
}
