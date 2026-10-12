package option

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	soil "github.com/funtimecoding/soil/pkg/constant"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"slices"
	"strings"
)

func New(
	raw string,
	verbose bool,
) *Lint {
	result := &Lint{Verbose: verbose}

	if raw != "" {
		for _, skip := range split.Comma(raw) {
			if !strings.Contains(skip, stringsConstant.Dot) &&
				!strings.HasSuffix(skip, stringsConstant.Slash) {
				skip = join.Empty(skip, stringsConstant.Slash)
			}

			result.Skips = append(result.Skips, skip)
		}
	}

	for _, skip := range []string{
		gitConstant.Directory,
		systemConstant.IdeaPath,
		systemConstant.FixturePath,
		systemConstant.Temporary,
	} {
		withSlash := fmt.Sprintf("%s/", skip)

		if !slices.Contains(result.Skips, withSlash) {
			result.Skips = append(result.Skips, withSlash)
		}
	}

	if !slices.Contains(result.Skips, soil.GeneratedFile) {
		result.Skips = append(result.Skips, soil.GeneratedFile)
	}

	if verbose {
		console.Format(
			"Skips (%d): %s\n",
			len(result.Skips),
			join.Comma(result.Skips),
		)
	}

	return result
}
