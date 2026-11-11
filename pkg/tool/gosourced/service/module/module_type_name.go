package module

import (
	"github.com/funtimecoding/soil/pkg/source/canonical_type"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/types"
	"strings"
)

func moduleTypeName(
	subject types.Type,
	modulePath string,
	newModulePath string,
) string {
	return canonical_type.String(
		subject,
		func(p *types.Package) string {
			for _, root := range []string{modulePath, newModulePath} {
				if p.Path() == root ||
					strings.HasPrefix(p.Path(), join.Empty(root, "/")) {
					return join.Empty(
						constant.ModuleQualifier,
						strings.TrimPrefix(p.Path(), root),
					)
				}
			}

			return p.Path()
		},
	)
}
