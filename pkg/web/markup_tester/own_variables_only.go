package markup_tester

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/style_sheet"
	"strings"
)

func ownVariablesOnly(r *style_sheet.Rule) bool {
	for _, d := range r.Declarations {
		if !strings.HasPrefix(d, constant.VariablePrefix) ||
			strings.HasPrefix(d, constant.PicoVariablePrefix) {
			return false
		}
	}

	return len(r.Declarations) > 0
}
