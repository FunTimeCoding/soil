package web

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/roster_card"

func (s *Server) rosterCards() ([]*roster_card.Card, error) {
	sessions, e := s.roster.ListSessions()

	if e != nil {
		return nil, e
	}

	identifiers := make([]string, 0, len(sessions))

	for _, v := range sessions {
		identifiers = append(identifiers, v.Identifier)
	}

	labels, f := s.label.LabelsBySessions(identifiers)

	if f != nil {
		return nil, f
	}

	pulses, g := s.message.LatestPulsesBySessions(identifiers)

	if g != nil {
		return nil, g
	}

	queued, h := s.message.QueuedCounts(sessions)

	if h != nil {
		return nil, h
	}

	cruise, i := s.message.CruiseSettings(identifiers)

	if i != nil {
		return nil, i
	}

	result := make([]*roster_card.Card, 0, len(sessions))

	for j := range sessions {
		c := roster_card.New(&sessions[j])
		c.Labels = labels[c.Session.Identifier]
		c.Latest = pulses[c.Session.Identifier]
		c.Queued = queued[c.Session.Identifier]
		c.Cruise = cruise[c.Session.Identifier]
		c.CruiseGrace = s.message.Grace(c.Cruise)
		result = append(result, c)
	}

	return result, nil
}
