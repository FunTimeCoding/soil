package installed

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/parser"
	"go/token"
	"path/filepath"
	"strconv"
	"strings"
)

func implementer(
	sources map[string]string,
	module string,
	name string,
) string {
	f, e := parser.ParseFile(
		token.NewFileSet(),
		filepath.Join(
			sources[module],
			constant.CommandDirectory,
			name,
			soil.MainFile,
		),
		nil,
		parser.ImportsOnly,
	)

	if e != nil {
		return module
	}

	result := ""

	for _, i := range f.Imports {
		path, g := strconv.Unquote(i.Path.Value)

		if g != nil {
			continue
		}

		for m := range sources {
			if path != m &&
				!strings.HasPrefix(path, join.Empty(m, stringsConstant.Slash)) {
				continue
			}

			if m == module {
				return module
			}

			if len(m) > len(result) {
				result = m
			}
		}
	}

	if result == "" {
		return module
	}

	return result
}
