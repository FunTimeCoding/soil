package scan

import (
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/parse"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
	"path/filepath"
	"strings"
)

func RegisteredTools(
	v *virtual_file_system.System,
	path string,
) ([]string, []*concern.Concern) {
	if !v.DirectoryExists(constant.ModelContextDirectory) {
		return nil, []*concern.Concern{
			concern.NewPackage(
				constant.MappedServiceMissingKey,
				constant.MappedServiceMissingText,
				path,
			),
		}
	}

	constants := constantStrings(v)
	var names []string
	var concerns []*concern.Concern
	prefix := fmt.Sprintf("%s/", constant.ModelContextDirectory)

	for _, file := range v.Files() {
		if !strings.HasPrefix(file, prefix) ||
			!strings.HasSuffix(file, library.GoExtension) {
			continue
		}

		f, _, e := parse.Source(filepath.Base(file), v.ReadString(file))

		if e != nil {
			concerns = append(
				concerns,
				concern.NewPackage(
					constant.FileNotParseableKey,
					constant.FileNotParseableText,
					filepath.Join(path, file),
				),
			)

			continue
		}

		for _, c := range parse.FindCalls(f, "mcp", "NewTool") {
			if len(c.Args) == 0 {
				continue
			}

			if literal, okay := parse.StringValue(c.Args[0]); okay {
				names = append(names, literal)

				continue
			}

			if selector, okay := c.Args[0].(*ast.SelectorExpr); okay {
				if value, found := constants[selector.Sel.Name]; found {
					names = append(names, value)

					continue
				}
			}

			concerns = append(
				concerns,
				concern.NewPackage(
					constant.UnresolvedToolNameKey,
					constant.UnresolvedToolNameText,
					filepath.Join(path, file),
				),
			)
		}
	}

	return names, concerns
}
