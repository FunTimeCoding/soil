package tracker_source

import "slices"

func (s *Source) Covers(release string) bool {
	return slices.Contains(s.releases, release)
}
