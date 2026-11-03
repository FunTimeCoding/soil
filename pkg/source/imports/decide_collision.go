package imports

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/types/import_decision"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
)

func decideCollision(
	v *import_view.View,
	declaredName string,
	group []*import_entry.Entry,
	chosen map[string]bool,
) []*import_decision.Decision {
	keeper, reason := collisionKeeper(v, declaredName, group)
	refused := ""

	for _, e := range group {
		if e.Path == keeper && e.Local != declaredName {
			refused = refusal(v, group, nil, declaredName)
		}
	}

	accepted := acceptedAliases(v, declaredName, keeper, group)

	for e := range accepted {
		chosen[e.Local] = true
	}

	var result []*import_decision.Decision

	for _, e := range group {
		d := import_decision.New(e, constant.RuleCollision)
		d.Keeper = keeper
		d.Reason = reason
		result = append(result, d)

		if e.Path == keeper {
			if e.Local == declaredName {
				continue
			}

			d.Want = declaredName

			if refused != "" {
				d.Kind = constant.DecisionRefused
				d.Reason = refused

				continue
			}

			d.Kind = constant.DecisionRename

			continue
		}

		if refused != "" || accepted[e] {
			d.Want = e.Local

			continue
		}

		want := loserName(v, declaredName, group, chosen, e)

		if want == "" {
			d.Kind = constant.DecisionUnderivable
			d.Subsystem = Subsystem(e.Path, e.Module)

			continue
		}

		chosen[want] = true
		d.Want = want

		if e.Local != want {
			d.Kind = constant.DecisionRename
		}
	}

	return result
}
