package symbol

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
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
	all, set, e := s.workspace.ReferenceLoad(
		directory,
		packagePath,
		"",
		func(loaded []*packages.Package, _ *token.FileSet) []types.Object {
			var result []types.Object

			for _, n := range renames {
				d, _, f := lookup.FindDeclaration(
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

	if renamesFamily(all, packagePath, renames) {
		all, set, e = s.workspace.Whole(directory, packagePath)

		if e != nil {
			return nil, e
		}
	}

	var planned []*plannedRename

	for _, n := range renames {
		d, _, f := lookup.FindDeclaration(
			all,
			packagePath,
			n.OldName,
			n.Receiver,
		)

		if f == nil && isInterfaceMethod(d) {
			planned = append(
				planned,
				s.planFamily(directory, all, set, d.(*types.Func), n, r)...,
			)

			continue
		}

		if p := s.checkRename(directory, all, set, packagePath, n, r); p != nil {
			planned = append(planned, p)
		}
	}

	if commit.HasUnfixed(r) {
		return r, nil
	}

	decorations := decoration.NewSet()

	for _, p := range planned {
		if f := renameReferences(decorations, all, set, p, r); f != nil {
			return nil, f
		}
	}

	e = decoration.Restore(decorations, resolve.NewNames(all), nil, out)

	if e != nil {
		return nil, e
	}

	for _, p := range planned {
		renameFollow(set, all, p.Declaration, p.Rename.NewName, out, r)
		declaring := filepath.Dir(set.Position(p.Declaration.Pos()).Filename)
		out.Resymbol(
			declaring,
			identifier.MemberName(p.Rename.Receiver, p.Rename.OldName),
			declaring,
			identifier.MemberName(p.Rename.Receiver, p.Rename.NewName),
		)
	}

	return r, nil
}
