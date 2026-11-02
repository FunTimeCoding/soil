package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol_rename"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func (s *Service) renameThrough(
	directory string,
	packagePath string,
	renames []*symbol_rename.Rename,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := s.referenceLoad(
		directory,
		packagePath,
		"",
		func(loaded []*packages.Package, _ *token.FileSet) []types.Object {
			var result []types.Object

			for _, n := range renames {
				d, _, f := findDeclaration(
					loaded,
					packagePath,
					n.OldName,
					n.Receiver,
				)

				if f == nil {
					result = append(result, d)
				}
			}

			return result
		},
	)

	if e != nil {
		return nil, e
	}

	var planned []*plannedRename

	for _, n := range renames {
		if p := s.checkRename(directory, all, set, packagePath, n, r); p != nil {
			planned = append(planned, p)
		}
	}

	if hasUnfixed(r) {
		return r, nil
	}

	decorations := decoration.NewSet()

	for _, p := range planned {
		if f := renameReferences(decorations, all, set, p, r); f != nil {
			return nil, f
		}
	}

	e = restoreDecorations(decorations, resolve.NewNames(all), nil, out)

	if e != nil {
		return nil, e
	}

	for _, p := range planned {
		renameFollow(set, all, p.Declaration, p.Rename.NewName, out, r)
		declaring := filepath.Dir(set.Position(p.Declaration.Pos()).Filename)
		out.Resymbol(
			declaring,
			memberName(p.Rename.Receiver, p.Rename.OldName),
			declaring,
			memberName(p.Rename.Receiver, p.Rename.NewName),
		)
	}

	return r, nil
}
