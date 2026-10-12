package imports

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/types/import_decision"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
)

func decideConstants(v *import_view.View) []*import_decision.Decision {
	var group []*import_entry.Entry
	var paths []string

	for _, e := range v.Entries {
		if e.Stays && e.Declared == constant.ConstantPackage {
			group = append(group, e)
			paths = append(paths, e.Path)
		}
	}

	keeper := ConstantKeeper(v.PackagePath, paths)
	chosen := make(map[string]bool)
	var result []*import_decision.Decision

	for _, e := range group {
		d := import_decision.New(e, constant.RuleConstant)
		d.Keeper = keeper
		want := constant.ConstantPackage

		if e.Path != keeper {
			want = ConstantAlias(e.Path, e.Module)
		}

		result = append(result, d)

		if want == "" {
			d.Kind = constant.DecisionUnderivable

			continue
		}

		d.Want = want

		if e.Local == want {
			chosen[want] = true

			continue
		}

		if r := refusal(v, group, chosen, want); r != "" {
			d.Kind = constant.DecisionRefused
			d.Reason = r

			continue
		}

		chosen[want] = true
		d.Kind = constant.DecisionRename
	}

	return result
}
