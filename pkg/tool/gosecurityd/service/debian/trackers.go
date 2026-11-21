package debian

import "github.com/funtimecoding/soil/pkg/vulnerability/debian/tracker"

func (s *Service) Trackers() []*tracker.Tracker {
	result := make([]*tracker.Tracker, len(s.sources))

	for i, source := range s.sources {
		result[i] = source.Tracker()
	}

	return result
}
