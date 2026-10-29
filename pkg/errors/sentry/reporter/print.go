package reporter

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/funtimecoding/soil/pkg/errors"
	"maps"
	"slices"
	"strings"
)

func (r *Reporter) print(
	v any,
	key string,
	context map[string]any,
) {
	var b strings.Builder
	b.WriteString(r.paint(color.FgRed, fmt.Sprintf("%s: %v", r.project, v)))
	b.WriteString("\n")

	for _, k := range slices.Sorted(maps.Keys(context)) {
		_, e := fmt.Fprintf(
			&b,
			"  %s %s %v\n",
			r.paint(color.FgCyan, key),
			k,
			context[k],
		)
		errors.PanicOnError(e)
	}

	all := frames(2)
	shown := 0

	for shown < len(all) && r.owned(all[shown].Function) {
		shown++
	}

	width := 0

	for _, f := range all[:shown] {
		width = max(width, len(r.location(f)))
	}

	for _, f := range all[:shown] {
		_, e := fmt.Fprintf(
			&b,
			"  %-*s  %s\n",
			width,
			r.location(f),
			r.paint(color.FgYellow, functionName(f.Function)),
		)
		errors.PanicOnError(e)
	}

	if rest := all[shown:]; len(rest) > 0 {
		b.WriteString(r.paint(color.Faint, r.collapsed(rest)))
	}

	_, e := r.output.Write([]byte(b.String()))
	errors.PanicOnError(e)
}
