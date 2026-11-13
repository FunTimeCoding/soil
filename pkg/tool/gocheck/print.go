package gocheck

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"os"
)

func Print(r *check_result.Result) bool {
	if r.Announcement != "" {
		writer.Print(os.Stderr, "%s\n", r.Announcement)
	}

	if r.Stale != "" {
		writer.Print(os.Stderr, constant.StaleNotice, r.Stale)
	}

	if r.Unchanged {
		console.Line(constant.NothingChanged)

		return false
	}

	blocked := false

	for _, s := range r.Sections {
		if output.Print(s, false) {
			blocked = true
		}
	}

	if len(r.Changes) > 0 {
		lint.Header(constant.Identity.Name(), r.Root, "")
		gofix.PrintChanges(r.Changes)
	}

	return blocked
}
