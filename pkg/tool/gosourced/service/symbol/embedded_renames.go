package symbol

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func embeddedRenames(
	all []*packages.Package,
	o types.Object,
	newName string,
	renames map[*ast.Ident]string,
) string {
	fields, capture := identifier.EmbeddedFieldReferences(all, o, newName)

	if capture != "" {
		return capture
	}

	for _, f := range fields {
		renames[f.Ident] = newName
	}

	return ""
}
