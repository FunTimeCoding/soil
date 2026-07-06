package option

import (
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	"slices"
)

func New(
	raw string,
	verbose bool,
) *Lint {
	result := &Lint{Raw: raw}

	if raw != "" {
		result.Skips = split.Comma(raw)
	}

	for _, skip := range []string{
		constant.Directory,
		system.IdeaPath,
		system.FixturePath,
		system.Temporary,
	} {
		withSlash := fmt.Sprintf("%s/", skip)

		if !slices.Contains(result.Skips, withSlash) {
			result.Skips = append(result.Skips, withSlash)
		}
	}

	if !slices.Contains(result.Skips, library.GeneratedFile) {
		result.Skips = append(result.Skips, library.GeneratedFile)
	}

	result.Count = len(result.Skips)

	if result.Count > 0 && verbose {
		fmt.Printf(
			"Skips (%d): %s\n",
			result.Count,
			join.Comma(result.Skips),
		)
	}

	return result
}
