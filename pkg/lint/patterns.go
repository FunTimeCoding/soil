package lint

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
	"strings"
)

func Patterns(
	root string,
	work string,
	patterns []string,
) ([]string, error) {
	var result []string
	current := join.Empty(soil.CurrentDirectory, stringsConstant.Slash)
	parent := join.Empty(soil.ParentDirectory, stringsConstant.Slash)

	for _, p := range patterns {
		if strings.HasSuffix(p, soil.GoExtension) {
			return nil, errors.Format(
				"provide a package directory pattern, not a file",
				p,
			)
		}

		if p != soil.CurrentDirectory &&
			!strings.HasPrefix(p, current) &&
			!strings.HasPrefix(p, parent) &&
			!filepath.IsAbs(p) {
			result = append(result, p)

			continue
		}

		path, recursive := strings.CutSuffix(p, go_modConstant.RecursivePattern)
		scopes, e := Scopes(root, work, []string{path})

		if e != nil {
			return nil, e
		}

		translated := soil.CurrentDirectory

		if len(scopes) > 0 {
			translated = join.Empty(current, scopes[0])
		}

		if recursive {
			translated = join.Empty(translated, go_modConstant.RecursivePattern)
		}

		result = append(result, translated)
	}

	return result, nil
}
