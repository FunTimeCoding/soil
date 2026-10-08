package goreplace

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/block"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/match"
	"io"
	"os"
	"strings"
)

func Run(
	path string,
	input string,
	dryRun bool,
	out io.Writer,
) error {
	blocks, e := block.Parse(input)

	if e != nil {
		return e
	}

	b, f := os.ReadFile(path)

	if f != nil {
		return f
	}

	content := string(b)
	matches, failures := match.Locate(content, blocks)

	if len(failures) > 0 {
		return fmt.Errorf(constant.Failed, path, strings.Join(failures, "\n"))
	}

	var result strings.Builder
	previous := 0

	for _, m := range matches {
		result.WriteString(content[previous:m.Offset])
		result.WriteString(m.Block.Replace)
		previous = m.Offset + len(m.Block.Search)
	}

	result.WriteString(content[previous:])
	printDifference(out, matches)

	if dryRun {
		write(out, constant.DryRunApplied, path, len(matches))

		return nil
	}

	system.WriteFile(path, []byte(result.String()), system.Mode(path))
	write(out, constant.Applied, path, len(matches))

	return nil
}
