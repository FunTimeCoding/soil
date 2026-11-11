package relocator

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func MethodEntries(
	set *token.FileSet,
	p *packages.Package,
	named *types.Named,
	methods []string,
	targetFile string,
) ([]*relocation.Entry, string) {
	var result []*relocation.Entry

	for _, name := range methods {
		var o *types.Func

		for m := range named.Methods() {
			if m.Name() == name {
				o = m
			}
		}

		if o == nil {
			return nil, fmt.Sprintf(
				"%s has no method %s declared on it",
				named.Obj().Name(),
				name,
			)
		}

		file, declaration, _ := findDeclarationNode(p, o)

		if declaration == nil {
			return nil, fmt.Sprintf("declaration not found: %s", name)
		}

		fileName := targetFile

		if fileName == "" {
			fileName = filepath.Base(set.Position(file.Pos()).Filename)
		}

		entry := relocation.NewEntry()
		entry.Symbol = name
		entry.NewName = name
		entry.Object = o
		entry.File = file
		entry.Declaration = declaration
		entry.Node = declaration
		entry.Carried = imports.UsedBy(file, declaration)
		entry.TargetFile = fileName
		result = append(result, entry)
	}

	return result, ""
}
