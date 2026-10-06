package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/references"
	"path/filepath"
)

func (s *Service) FileReferences(
	directory string,
	packagePath string,
	filePath string,
) (*output.Results, *result.FileReferences, error) {
	r := output.NewResultsWithDirectory(directory)
	pattern := packagePath

	if s.full {
		pattern = "./..."
	}

	all, set, e := loadPackages(directory, pattern)

	if e != nil {
		return nil, nil, e
	}

	p := findPackage(all, packagePath)

	if p == nil {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernValidation,
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
				constant.ConcernValidation,
				fmt.Sprintf("file not found in %s: %s", packagePath, filePath),
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
				constant.ConcernValidation,
				"file declares no top-level symbols",
				"",
				false,
			),
		)

		return r, nil, nil
	}

	var entries []*references.References
	var i *xref.Index

	if !s.full {
		i = s.workspace(directory).References()
	}

	for _, q := range symbols {
		declaration, _, f := findDeclaration(
			all,
			packagePath,
			q.Name,
			q.Receiver,
		)

		if f != nil {
			r.AddConcern(
				concern.NewFile(
					constant.ConcernValidation,
					f.Error(),
					"",
					false,
				),
			)

			return r, nil, nil
		}

		name := q.Name

		if q.Receiver != "" {
			name = join.Empty(q.Receiver, constant.MemberSeparator, q.Name)
		}

		locations := referenceLocations(directory, all, set, declaration, full)

		if i != nil {
			locations = indexedLocations(
				directory,
				i,
				all,
				set,
				packagePath,
				declaration,
				full,
			)
		}

		entries = append(entries, references.New(name, locations))
	}

	return r, result.NewFileReferences(filePath, entries), nil
}
