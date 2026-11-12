package gofix

import (
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"github.com/funtimecoding/soil/pkg/strings/camel"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"strings"
)

func identityMoves(
	fileSet *token.FileSet,
	all []*packages.Package,
	violations []Violation,
	w *workspace.Workspace,
	r *output.Results,
) {
	files := syntaxFiles(fileSet, all)

	for _, v := range violations {
		switch v.object.(type) {
		case *types.TypeName, *types.Func:
		default:
			continue
		}

		path := fileSet.File(v.ident.Pos()).Name()

		if v.fix == "" || v.object.Pos() != v.ident.Pos() ||
			strings.HasSuffix(path, soil.TestSuffix) ||
			files[path] == nil || soleIdentity(files[path]) != v.ident.Name {
			continue
		}

		stem := strings.TrimSuffix(filepath.Base(path), soil.GoExtension)
		suffix := platformSuffix(stem)

		if strings.TrimSuffix(stem, suffix) != camel.ToSnake(v.ident.Name) {
			continue
		}

		name := join.Empty(
			camel.ToSnake(
				segment.ReplaceSegment(v.ident.Name, v.segment, v.fix),
			),
			suffix,
			soil.GoExtension,
		)
		target := filepath.Join(filepath.Dir(path), name)

		if target == path {
			continue
		}

		if w.Taken(target) {
			r.AddConcern(
				concern.NewFile(
					"file_identity",
					fmt.Sprintf(
						"%s is taken; %s keeps its name",
						name,
						filepath.Base(path),
					),
					path,
					false,
				),
			)

			continue
		}

		w.Move(path, target)
		r.AddConcern(
			concern.NewFile(
				"moved",
				fmt.Sprintf("moved %s → %s", filepath.Base(path), name),
				path,
				true,
			),
		)
	}
}
