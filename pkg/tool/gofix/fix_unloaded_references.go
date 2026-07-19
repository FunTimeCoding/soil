package gofix

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"go/ast"
	"os"
	"path/filepath"
	"strings"
)

func FixUnloadedReferences(
	violations []violation,
	loadedFiles map[string]bool,
	directory string,
	r *output.Results,
) {
	var renames []exportedRename

	for _, v := range violations {
		if v.fix == "" {
			continue
		}

		if !ast.IsExported(v.ident.Name) {
			continue
		}

		renames = append(
			renames,
			exportedRename{
				oldName: v.ident.Name,
				newName: segment.ReplaceSegment(
					v.ident.Name,
					v.segment,
					v.fix,
				),
			},
		)
	}

	if len(renames) == 0 {
		return
	}

	root := directory

	if root == "" {
		root = "."
	}

	errors.PanicOnError(
		filepath.Walk(
			root,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return nil
				}

				if i.IsDir() {
					return nil
				}

				if !strings.HasSuffix(path, constant.GoExtension) {
					return nil
				}

				full, e := filepath.Abs(path)
				errors.PanicOnError(e)

				if loadedFiles[full] {
					return nil
				}

				fixFileReferences(path, renames, r)

				return nil
			},
		),
	)
}
