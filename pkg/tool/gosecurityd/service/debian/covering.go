package debian

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/tracker_source"

func (s *Service) Covering(release string) *tracker_source.Source {
	for _, source := range s.sources {
		if source.Covers(release) {
			return source
		}
	}

	return nil
}
