package gobreak

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_result"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"os"
)

func report(
	results []*break_result.Result,
	passed int,
	asNotation bool,
) {
	missed := 0

	for _, r := range results {
		if r.Missed() {
			missed++
		}
	}

	if asNotation {
		writer.Print(
			os.Stdout,
			constant.LinePrint,
			notation.Encode(results, false),
		)
	} else {
		writer.Print(os.Stdout, constant.BaselinePassed, passed)

		for _, r := range results {
			writer.Print(os.Stdout, constant.LinePrint, r.Line())
		}

		writer.Print(
			os.Stdout,
			constant.Summary,
			len(results),
			len(results)-missed,
		)
	}

	if missed > 0 {
		system.Exitf(1, constant.MissedFormat, missed)
	}
}
