package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"path/filepath"
)

func (s *Service) FileReferences(
	directory string,
	packagePath string,
	filePath string,
) (*output.Results, *result.FileReferences, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := resolve.LoadPackages(directory, "./...")

	if e != nil {
		return nil, nil, e
	}

	p := findPackage(all, packagePath)

	if p == nil {
		r.AddConcern(
			concern.NewFile(
				"validation",
				fmt.Sprintf("package not found: %s", packagePath),
				"",
				false,
			),
		)

		return r, nil, nil
	}

	full := filePath

	if !filepath.IsAbs(full) {
		full = filepath.Join(directory, filePath)
	}

	file := findSyntaxFile(set, p, full)

	if file == nil {
		r.AddConcern(
			concern.NewFile(
				"validation",
				fmt.Sprintf(
					"file not found in %s: %s",
					packagePath,
					filePath,
				),
				"",
				false,
			),
		)

		return r, nil, nil
	}

	symbols := expandFileQuerySymbols(file)

	if len(symbols) == 0 {
		r.AddConcern(
			concern.NewFile(
				"validation",
				"file declares no top-level symbols",
				"",
				false,
			),
		)

		return r, nil, nil
	}

	var entries []*result.References

	for _, q := range symbols {
		declaration, _, f := findDeclaration(
			all,
			packagePath,
			q.name,
			q.receiver,
		)

		if f != nil {
			r.AddConcern(
				concern.NewFile("validation", f.Error(), "", false),
			)

			return r, nil, nil
		}

		name := q.name

		if q.receiver != "" {
			name = join.Empty(
				q.receiver,
				constant.MemberSeparator,
				q.name,
			)
		}

		entries = append(
			entries,
			result.NewReferences(
				name,
				referenceLocations(directory, all, set, declaration, full),
			),
		)
	}

	return r, result.NewFileReferences(filePath, entries), nil
}
