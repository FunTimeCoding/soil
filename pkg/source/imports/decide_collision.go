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
) []*import_decision.Decision {
	var paths []string
	plain := ""
	standard := make(map[string]bool)

	for _, e := range group {
		paths = append(paths, e.Path)
		standard[e.Path] = e.Standard

		if e.Local == declaredName {
			plain = e.Path
		}
	}

	keeper, reason := Keeper(
		v.PackagePath,
		declaredName,
		paths,
		plain,
		func(p string) bool { return standard[p] },
	)
	chosen := make(map[string]bool)
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

			if r := refusal(v, group, nil, declaredName); r != "" {
				d.Kind = constant.DecisionRefused
				d.Reason = r

				continue
			}

			d.Kind = constant.DecisionRename

			continue
		}

		want := ChooseAlias(
			e.Path,
			e.Module,
			declaredName,
			func(c string) bool {
				return c == declaredName || refusal(v, group, chosen, c) != ""
			},
		)

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
