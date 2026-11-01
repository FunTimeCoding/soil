package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"go/token"
	"maps"
	"path/filepath"
	"slices"
)

func reportClauses(
	r *output.Results,
	set *token.FileSet,
	modified map[string]*ast.File,
	directory string,
	oldName string,
	newName string,
) {
	testOld := join.Empty(oldName, constant.TestClauseSuffix)

	for _, filename := range slices.Sorted(maps.Keys(modified)) {
		file := modified[filename]
		to := newName

		if file.Name.Name == testOld {
			to = join.Empty(newName, constant.TestClauseSuffix)
		}

		r.AddConcern(
			concern.NewLine(
				constant.ConcernRenamed,
				fmt.Sprintf("package %s → %s", file.Name.Name, to),
				filepath.Join(directory, filepath.Base(filename)),
				set.Position(file.Name.Pos()).Line,
				"",
				true,
			),
		)
	}
}
