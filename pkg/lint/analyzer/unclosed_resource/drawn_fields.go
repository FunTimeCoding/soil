package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"maps"
	"slices"
)

func (s *Summaries) drawnFields(w fact.FieldForward) []string {
	var forwarded *fact.Condition

	for i := range w.Conditions {
		c := &w.Conditions[i]

		if c.Kind == constant.DrawsCondition && c.Field == constant.AnyField {
			forwarded = c

			continue
		}

		if !s.holds(c) {
			return nil
		}
	}

	if forwarded == nil {
		return []string{w.Field}
	}

	return slices.Collect(
		maps.Keys(s.draws[forwarded.Callee][forwarded.Argument]),
	)
}
