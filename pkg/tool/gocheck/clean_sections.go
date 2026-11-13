package gocheck

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/output"
	goanalyzeConstant "github.com/funtimecoding/soil/pkg/tool/goanalyze/constant"
	gofixConstant "github.com/funtimecoding/soil/pkg/tool/gofix/constant"
)

func cleanSections(
	root string,
	patterns []string,
) []*output.Results {
	var result []*output.Results

	for _, name := range []string{
		gofixConstant.Identity.Name(),
		goanalyzeConstant.Identity.Name(),
	} {
		r := output.NewResultsWithDirectory(root)
		r.Header = lint.HeaderLine(name, root, lint.PatternDetail(patterns))
		result = append(result, r)
	}

	return result
}
