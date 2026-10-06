package sink

import "slices"

func (s *Sink) Dropped() []string {
	return slices.Clone(s.dropped)
}
