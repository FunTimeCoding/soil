package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
	"golang.org/x/tools/go/packages"
)

func AnalyzeLoaded(
	o *option.Analyze,
	raw []*packages.Package,
	reported map[string]bool,
	handle *index.Workspace,
) *output.Results {
	if !whole(o, analyzedPatterns(o)) {
		loaded, faces, summaries := persistedFrom(raw, handle)

		return report(o, loaded, reported, faces, summaries)
	}

	loaded, faces, summaries := fullIndex(o.Root, raw, handle, o.Notice)

	return report(o, loaded, reported, faces, summaries)
}
