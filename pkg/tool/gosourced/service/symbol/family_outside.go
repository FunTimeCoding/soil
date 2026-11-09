package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func familyOutside(
	interfaces *face.Set,
	all []*packages.Package,
	members []*types.Func,
) []string {
	modules := workspaceModules(all)
	var result []string

	for _, f := range members {
		if isInterfaceMethod(f) {
			continue
		}

		for _, c := range interfaces.Satisfied(f) {
			if _, declares := c.Methods[f.Name()]; !declares ||
				inWorkspace(modules, c.Package) {
				continue
			}

			result = append(
				result,
				fmt.Sprintf(
					"%s satisfies %s.%s outside the workspace - it can't be renamed with the family",
					types.TypeString(receiverNamed(f), nil),
					c.Package,
					c.Name,
				),
			)
		}
	}

	return result
}
