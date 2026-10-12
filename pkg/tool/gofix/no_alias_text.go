package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/source/types/import_decision"
)

func noAliasText(d *import_decision.Decision) string {
	e := d.Entry

	if d.Rule == constant.RuleConstant || d.Subsystem == "" {
		return fmt.Sprintf(
			"no subsystem alias for %s, keeping %s",
			e.Path,
			e.Local,
		)
	}

	return fmt.Sprintf(
		"no free alias for %s (%s and %s taken), keeping %s",
		e.Path,
		d.Subsystem,
		imports.Fallback(d.Subsystem, e.Declared),
		e.Local,
	)
}
