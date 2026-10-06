package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
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
) ([]*relocation.Entry, string) {
	named, okay := typeObject.Type().(*types.Named)

	if !okay {
		return nil, fmt.Sprintf("%s is not a named type", typeObject.Name())
	}

	objects := []types.Object{typeObject}

	for i := range named.NumMethods() {
		objects = append(objects, named.Method(i))
	}

	var result []*relocation.Entry

	for _, o := range objects {
		file, declaration, spec := findDeclarationNode(p, o)

		if declaration == nil {
			return nil, fmt.Sprintf("declaration not found: %s", o.Name())
		}

		node := ast.Node(declaration)

		if spec != nil {
			node = spec
		}

		name := targetFile

		if name == "" {
			name = filepath.Base(set.Position(file.Pos()).Filename)
		}

		entry := relocation.NewEntry()
		entry.Symbol = o.Name()
		entry.NewName = o.Name()
		entry.Object = o
		entry.File = file
		entry.Declaration = declaration
		entry.Spec = spec
		entry.Node = node
		entry.Carried = imports.UsedBy(file, node)
		entry.TargetFile = name
		result = append(result, entry)
	}

	return result, ""
}
