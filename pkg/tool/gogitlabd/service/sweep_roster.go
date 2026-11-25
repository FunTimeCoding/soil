package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/prometheus/common/model"
)

func (s *Service) SweepRoster(x context.Context) []error {
	if s.store == nil {
		return nil
	}

	window := model.Duration(constant.RosterWindow).String()
	roster, e := peer.Guarded(
		x,
		s.rosterGuard,
		func(y context.Context) (*client.RosterResponse, error) {
			return s.roster.GetRoster(
				y,
				&client.GetRosterParams{SeenWithin: &window},
			)
		},
	)

	if e != nil {
		if refused := s.note(constant.RosterPeer, e); refused != nil {
			return []error{refused}
		}

		return nil
	}

	present := make([]string, 0, len(roster.Entries))

	for _, v := range roster.Entries {
		present = append(present, v.Callsign)
	}

	dropped, f := s.store.Abandon(present)

	if f != nil {
		return []error{f}
	}

	var result []error
	seen := map[string]bool{}

	for _, v := range dropped {
		if k := v.Key(); !seen[k.String()] {
			seen[k.String()] = true

			if g := s.forget(k); g != nil {
				result = append(result, g)
			}
		}
	}

	if len(dropped) > 0 {
		s.logger.Structured("subscription_abandon", "count", len(dropped))
	}

	return result
}
