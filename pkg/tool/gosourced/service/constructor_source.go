package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/format"
	"go/types"
	"sort"
	"strings"
)

func constructorSource(
	home *types.Package,
	name string,
	constructor string,
	parameters []*types.Var,
) ([]byte, error) {
	imports := map[string]string{}
	var conflict error
	qualifier := func(other *types.Package) string {
		if other == home {
			return ""
		}

		for path, known := range imports {
			if known == other.Name() && path != other.Path() {
				conflict = validation.New(
					"parameter types import two packages named %s",
					known,
				)
			}
		}

		imports[other.Path()] = other.Name()

		return other.Name()
	}
	var signature, assignments []string

	for _, v := range parameters {
		parameter := parameterName(v.Name())
		signature = append(
			signature,
			join.Space(parameter, types.TypeString(v.Type(), qualifier)),
		)
		assignments = append(
			assignments,
			fmt.Sprintf("%s: %s", v.Name(), parameter),
		)
	}

	if conflict != nil {
		return nil, conflict
	}

	var paths []string

	for path := range imports {
		paths = append(paths, fmt.Sprintf("%q", path))
	}

	sort.Strings(paths)
	parts := []string{fmt.Sprintf("package %s\n", home.Name())}

	switch len(paths) {
	case 0:
	case 1:
		parts = append(parts, fmt.Sprintf("import %s\n", paths[0]))
	default:
		parts = append(
			parts,
			fmt.Sprintf("import (\n\t%s\n)\n", strings.Join(paths, "\n\t")),
		)
	}

	list := ""

	switch len(signature) {
	case 0:
	case 1:
		list = signature[0]
	default:
		list = fmt.Sprintf("\n\t%s,\n", strings.Join(signature, ",\n\t"))
	}

	parts = append(
		parts,
		fmt.Sprintf(
			"func %s(%s) *%s {\n\treturn &%s{%s}\n}\n",
			constructor,
			list,
			name,
			name,
			join.CommaSpace(assignments),
		),
	)

	return format.Source([]byte(strings.Join(parts, "\n")))
}
