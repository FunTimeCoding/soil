package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/source/index"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func familyOutside(
	interfaces *face.Set,
	w *index.Workspace,
	all []*packages.Package,
	members []*types.Func,
) ([]string, []string) {
	modules := workspaceModules(all)
	wanted := map[*types.Named]bool{}

	for _, f := range members {
		wanted[receiverNamed(f).Origin()] = true
	}

	holders := holderPackages(all, modules, wanted)
	var refusals []string
	var notes []string

	for _, f := range members {
		if isInterfaceMethod(f) {
			continue
		}

		for _, c := range interfaces.Satisfied(f) {
			if _, declares := c.Methods[f.Name()]; !declares ||
				inWorkspace(modules, c.Package) {
				continue
			}

			name := types.TypeString(receiverNamed(f), nil)

			if holder, reaching := reachingHolder(w, holders, c.Package); reaching {
				refusals = append(
					refusals,
					fmt.Sprintf(
						"%s satisfies %s.%s outside the workspace, and %s reaches %s - it can't be renamed with the family",
						name,
						c.Package,
						c.Name,
						holder,
						c.Package,
					),
				)

				continue
			}

			notes = append(
				notes,
				fmt.Sprintf(
					"%s satisfies %s.%s outside the workspace by shape only - no package holding it reaches %s, so the rename drops that satisfaction",
					name,
					c.Package,
					c.Name,
					c.Package,
				),
			)
		}
	}

	return refusals, notes
}
