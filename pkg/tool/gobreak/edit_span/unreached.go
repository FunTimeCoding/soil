package edit_span

import "golang.org/x/tools/cover"

func (s *Span) Unreached(profiles []*cover.Profile) bool {
	overlapped := false

	for _, p := range profiles {
		if p.FileName != s.Profile {
			continue
		}

		for _, b := range p.Blocks {
			if b.EndLine < s.First || b.StartLine > s.Last {
				continue
			}

			if b.Count > 0 {
				return false
			}

			overlapped = true
		}
	}

	return overlapped
}
