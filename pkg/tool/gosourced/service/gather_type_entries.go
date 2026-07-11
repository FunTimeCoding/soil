package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func gatherTypeEntries(
	set *token.FileSet,
	p *packages.Package,
	typeObject types.Object,
	targetFile string,
) ([]*moveEntry, string) {
	named, okay := typeObject.Type().(*types.Named)

	if !okay {
		return nil, fmt.Sprintf(
			"%s is not a named type",
			typeObject.Name(),
		)
	}

	objects := []types.Object{typeObject}

	for i := range named.NumMethods() {
		objects = append(objects, named.Method(i))
	}

	var result []*moveEntry

	for _, o := range objects {
		file, declaration, spec := findDeclarationNode(p, o)

		if declaration == nil {
			return nil, fmt.Sprintf(
				"declaration not found: %s",
				o.Name(),
			)
		}

		node := ast.Node(declaration)

		if spec != nil {
			node = spec
		}

		name := targetFile

		if name == "" {
			name = filepath.Base(
				set.Position(file.Pos()).Filename,
			)
		}

		result = append(
			result,
			&moveEntry{
				symbol:      o.Name(),
				newName:     o.Name(),
				object:      o,
				file:        file,
				declaration: declaration,
				spec:        spec,
				node:        node,
				carried:     imports.UsedBy(file, node),
				targetFile:  name,
			},
		)
	}

	return result, ""
}
