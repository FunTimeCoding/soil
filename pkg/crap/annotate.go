package crap

import (
	"github.com/funtimecoding/soil/pkg/crap/mutation"
	"github.com/funtimecoding/soil/pkg/crap/report"
	"path"
	"path/filepath"
	"strings"
)

func Annotate(
	r *report.Report,
	m *mutation.Report,
) {
	byFile := m.ByFile()

	for _, e := range r.Entries {
		if e.Missing || e.Coverage == 0 {
			continue
		}

		relative := filepath.ToSlash(
			strings.TrimPrefix(e.Function.File, r.Root),
		)
		mutants := byFile[path.Join(m.Module, strings.TrimPrefix(relative, "/"))]

		if len(mutants) == 0 {
			continue
		}

		e.Annotate(
			mutation.Within(mutants, e.Function.Line, e.Function.EndLine),
		)
	}
}
