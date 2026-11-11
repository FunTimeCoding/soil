package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"go/types"
	"golang.org/x/tools/go/packages"
	"slices"
)

func (s *Service) checkMovedInterfaces(
	directory string,
	all []*packages.Package,
	entries []*relocation.Entry,
	typeName string,
) string {
	faces := face.FromWorkspace(s.workspace.Workspace(directory), all)

	for _, entry := range entries {
		f := entry.Object.(*types.Func)
		var names []string

		for _, i := range faces.Satisfied(f) {
			names = append(names, join.Dot([]string{i.Package, i.Name}))
		}

		if len(names) > 0 {
			return fmt.Sprintf(
				"%s satisfies %s - moving it off %s breaks that interface",
				f.Name(),
				join.CommaSpace(
					slices.Compact(slices.Sorted(slices.Values(names))),
				),
				typeName,
			)
		}
	}

	return ""
}
