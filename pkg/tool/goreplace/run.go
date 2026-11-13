package goreplace

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/text/literal_edit/block"
	"github.com/funtimecoding/soil/pkg/text/literal_edit/match"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
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

	printDifference(out, content, matches)

	if dryRun {
		write(out, constant.DryRunApplied, path, len(matches))

		return nil
	}

	system.WriteFile(
		path,
		[]byte(match.Apply(content, matches)),
		system.Mode(path),
	)
	write(out, constant.Applied, path, len(matches))

	return nil
}
