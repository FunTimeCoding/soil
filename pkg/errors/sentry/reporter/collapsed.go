package reporter

import (
	"fmt"
	"runtime"
	"slices"
	"strings"
)

func (r *Reporter) collapsed(rest []runtime.Frame) string {
	var labels []string

	for _, f := range rest {
		l := r.label(f.Function)

		if !slices.Contains(labels, l) {
			labels = append(labels, l)
		}
	}

	if len(labels) > 3 {
		labels = append(labels[:3], "…")
	}

	noun := "frames"

	if len(rest) == 1 {
		noun = "frame"
	}

	return fmt.Sprintf(
		"  · %d more %s: %s\n",
		len(rest),
		noun,
		strings.Join(labels, ", "),
	)
}
