package lint

import (
	"github.com/funtimecoding/soil/pkg/constant"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
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
	current := join.Empty(constant.CurrentDirectory, stringsConstant.Slash)
	parent := join.Empty(constant.ParentDirectory, stringsConstant.Slash)

	for _, p := range patterns {
		if p != constant.CurrentDirectory &&
			!strings.HasPrefix(p, current) &&
			!strings.HasPrefix(p, parent) &&
			!filepath.IsAbs(p) {
			result = append(result, p)

			continue
		}

		path, recursive := strings.CutSuffix(p, lintConstant.RecursivePattern)
		scopes, e := Scopes(root, work, []string{path})

		if e != nil {
			return nil, e
		}

		translated := constant.CurrentDirectory

		if len(scopes) > 0 {
			translated = join.Empty(current, scopes[0])
		}

		if recursive {
			translated = join.Empty(translated, lintConstant.RecursivePattern)
		}

		result = append(result, translated)
	}

	return result, nil
}
