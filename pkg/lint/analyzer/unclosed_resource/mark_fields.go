package unclosed_resource

import "github.com/funtimecoding/soil/pkg/lint/fact"

func (s *Summaries) markFields(
	yielding map[string][]fact.FieldForward,
	drawing map[string][]fact.FieldForward,
) bool {
	changed := false

	for f, all := range yielding {
		for _, w := range all {
			if s.holdsAll(w.Conditions) &&
				markYield(s.yields, f, w.Argument, w.Field) {
				changed = true
			}
		}
	}

	for f, all := range drawing {
		for _, w := range all {
			for _, field := range s.drawnFields(w) {
				if markYield(s.draws, f, w.Argument, field) {
					changed = true
				}
			}
		}
	}

	return changed
}
