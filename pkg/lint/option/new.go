package option

import (
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/lint/document_cache"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"os"
	"slices"
	"strings"
)

func New(
	raw string,
	verbose bool,
) *Lint {
	result := &Lint{
		Verbose:   verbose,
		Trace:     os.Stdout,
		Trees:     warm_forest.New(),
		Documents: document_cache.New(),
	}

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
		writer.Print(
			result.Trace,
			"Skips (%d): %s\n",
			len(result.Skips),
			join.Comma(result.Skips),
		)
	}

	return result
}
