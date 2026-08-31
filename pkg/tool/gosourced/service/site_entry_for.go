package service

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"go/ast"
	"go/token"
	"os"
)

func (s *Service) siteEntryFor(
	directory string,
	set *token.FileSet,
	contents map[string][]byte,
	node ast.Node,
	anchor ast.Node,
	reference resolve.Reference,
) (*siteEntry, error) {
	position := set.Position(reference.Ident.Pos())
	content, okay := contents[position.Filename]

	if !okay {
		read, e := os.ReadFile(position.Filename)

		if e != nil {
			return nil, e
		}

		content = read
		contents[position.Filename] = read
	}

	shape, exemplar := statementShape(content, set, node, anchor)

	return &siteEntry{
		shape:    shape,
		exemplar: exemplar,
		location: result.NewLocation(
			system.RelativePath(directory, position.Filename),
			position.Line,
			reference.Package.PkgPath,
		),
	}, nil
}
