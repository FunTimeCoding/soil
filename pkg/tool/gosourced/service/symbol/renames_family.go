package symbol

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol_rename"
	"golang.org/x/tools/go/packages"
)

func renamesFamily(
	all []*packages.Package,
	packagePath string,
	renames []*symbol_rename.Rename,
) bool {
	for _, n := range renames {
		d, _, e := lookup.FindDeclaration(
			all,
			packagePath,
			n.OldName,
			n.Receiver,
		)

		if e == nil && isInterfaceMethod(d) {
			return true
		}
	}

	return false
}
