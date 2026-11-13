package goanalyze

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
)

func analyzedPatterns(o *option.Analyze) []string {
	if len(o.Patterns) == 0 {
		return []string{goModConstant.AllPackages}
	}

	return o.Patterns
}
