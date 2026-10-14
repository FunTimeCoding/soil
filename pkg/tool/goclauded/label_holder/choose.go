package label_holder

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"strings"
	"time"
)

func Choose(
	targets []*target.Target,
	r *Rule,
	now time.Time,
) *Plan {
	result := &Plan{}

	for _, t := range targets {
		value := t.Labels[r.Key]

		if !strings.EqualFold(value, r.Value) {
			continue
		}

		if Active(t, now, r.IdleThreshold) && result.Holder == "" {
			result.Holder = t.Identifier

			continue
		}

		result.Unset = append(
			result.Unset,
			&Release{Identifier: t.Identifier, Value: value},
		)
	}

	if result.Holder != "" {
		return result
	}

	for _, t := range targets {
		if !Active(t, now, r.IdleThreshold) || t.Labels[r.Key] != "" {
			continue
		}

		if t.TurnCount < r.MinimumTurns {
			continue
		}

		result.Adopt = t.Identifier
		result.Holder = t.Identifier

		break
	}

	return result
}
