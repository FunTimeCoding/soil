package service

import (
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern_site"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/location"
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
	reference resolve_reference.Reference,
) (*pattern_site.Entry, error) {
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

	return pattern_site.NewEntry(
		shape,
		exemplar,
		location.New(
			system.RelativePath(directory, position.Filename),
			position.Line,
			reference.Package.PkgPath,
		),
	), nil
}
