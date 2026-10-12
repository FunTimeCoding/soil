package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/types/import_decision"
)

func reportDecision(
	d *import_decision.Decision,
	filePath string,
	realias bool,
	r *output.Results,
) {
	e := d.Entry
	deAlias := d.Want == e.Declared
	verb := "alias"
	done := "re-aliased"

	if deAlias {
		verb = "de-alias"
		done = "de-aliased"
	}

	switch d.Kind {
	case constant.DecisionUnderivable:
		addSuggestion(r, filePath, noAliasText(d))
	case constant.DecisionRefused:
		addSuggestion(
			r,
			filePath,
			fmt.Sprintf(
				"cannot %s %s → %s (%s)",
				verb,
				e.Local,
				d.Want,
				d.Reason,
			),
		)
	case constant.DecisionRename:
		if d.Rule == constant.RuleConstant {
			reportAlias(
				r,
				filePath,
				fmt.Sprintf(
					"suggest %s %s → %s (%s)",
					verb,
					e.Local,
					d.Want,
					e.Path,
				),
				fmt.Sprintf("%s %s → %s (%s)", done, e.Local, d.Want, e.Path),
				realias,
			)

			return
		}

		if deAlias {
			addSuggestion(
				r,
				filePath,
				fmt.Sprintf(
					"suggest de-alias %s → %s (%s: %s)",
					e.Local,
					d.Want,
					d.Reason,
					e.Path,
				),
			)

			return
		}

		addSuggestion(
			r,
			filePath,
			fmt.Sprintf(
				"suggest alias %s → %s (%s; %s)",
				e.Local,
				d.Want,
				e.Path,
				keeperReason(d.Keeper, d.Reason),
			),
		)
	}
}
