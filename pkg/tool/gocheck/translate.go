package gocheck

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/types/layer_change"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/text/line_diff"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/constant"
	"path/filepath"
)

func translate(
	r *output.Results,
	root string,
	changes []*layer_change.Change,
) {
	anchors := map[string][]*line_diff.Anchor{}
	paths := map[string]string{}

	for _, c := range changes {
		if c.Current == nil {
			continue
		}

		destination, e := filepath.Rel(root, c.Destination)
		errors.PanicOnError(e)
		path, f := filepath.Rel(root, c.Path)
		errors.PanicOnError(f)
		anchors[destination] = line_diff.Anchors(
			string(c.Original),
			string(c.Current),
		)
		paths[destination] = path
	}

	for _, c := range r.Entries {
		lines, changed := anchors[c.Path]

		if !changed || c.Line < 1 || c.Line > len(lines) {
			continue
		}

		a := lines[c.Line-1]
		c.Path = paths[c.Path]
		c.Line = a.Line

		if a.Rewritten {
			c.Text = join.Empty(c.Text, constant.RewrittenSuffix)
		}
	}
}
