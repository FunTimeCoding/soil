package alert

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/maps"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (a *Alert) formatRemainingLabels(f *option.Format) string {
	if len(a.Remaining) == 0 {
		return ""
	}

	var result []string

	for _, k := range maps.StringKeys(a.Remaining) {
		v := a.Remaining[k]

		if f.UseColor {
			k = console.Yellow("%s", k)
		}

		result = append(result, fmt.Sprintf("%s=%s", k, v))
	}

	return join.Comma(result)
}
