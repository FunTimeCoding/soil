package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"go/format"
	"os"
	"strings"
)

func writeCreatedFile(
	packageName string,
	carried []*ast.ImportSpec,
	rendered string,
	path string,
) error {
	var builder strings.Builder
	builder.WriteString(join.Empty("package ", packageName, "\n\n"))

	if len(carried) > 0 {
		builder.WriteString("import (\n")

		for _, spec := range carried {
			builder.WriteString("\t")

			if spec.Name != nil {
				builder.WriteString(join.Empty(spec.Name.Name, " "))
			}

			builder.WriteString(join.Empty(spec.Path.Value, "\n"))
		}

		builder.WriteString(")\n\n")
	}

	builder.WriteString(join.Empty(rendered, "\n"))
	formatted, e := format.Source([]byte(builder.String()))

	if e != nil {
		return e
	}

	return os.WriteFile(path, formatted, 0644)
}
