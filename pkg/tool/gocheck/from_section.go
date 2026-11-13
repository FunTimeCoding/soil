package gocheck

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

func fromSection(
	root string,
	s client.CheckSection,
) *output.Results {
	result := output.NewResultsWithDirectory(root)
	result.Header = s.Header

	for _, c := range s.Entries {
		n := concern.New(c.Key, c.Text, c.Path, c.Line, c.LineText, c.Fixed)
		n.Type = c.Type
		n.Planned = c.Planned
		result.AddConcern(n)
	}

	for _, u := range s.Unchecked {
		result.AddUnchecked(
			u.Path,
			u.Line,
			u.Span,
			lintConstant.Reason(u.Reason),
		)
	}

	return result
}
