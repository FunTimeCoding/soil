package tracker_source

import "github.com/funtimecoding/soil/pkg/vulnerability/debian/tracker"

func (s *Source) Tracker() *tracker.Tracker {
	return s.tracker.Load()
}
