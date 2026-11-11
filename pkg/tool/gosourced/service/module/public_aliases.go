package module

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/token"
	"go/types"
	"strings"
)

func publicAliases(
	target *types.Package,
	modulePath string,
	newModulePath string,
) map[string]string {
	result := map[string]string{}

	for _, name := range target.Scope().Names() {
		o, okay := target.Scope().Lookup(name).(*types.TypeName)

		if !okay || !o.IsAlias() || !o.Exported() {
			continue
		}

		n, named := types.Unalias(o.Type()).(*types.Named)

		if !named || n.Obj().Pkg() == nil || n.Obj().Pkg() == target {
			continue
		}

		path := n.Obj().Pkg().Path()

		if path != newModulePath &&
			!strings.HasPrefix(path, join.Empty(newModulePath, "/")) {
			continue
		}

		public := types.NewNamed(
			types.NewTypeName(token.NoPos, target, o.Name(), nil),
			n.Underlying(),
			nil,
		)
		result[moduleTypeName(n, modulePath, newModulePath)] = moduleTypeName(
			public,
			modulePath,
			newModulePath,
		)
	}

	return result
}
