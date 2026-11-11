package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/types"
)

func checkTargetCollision(
	entries []*relocation.Entry,
	target *types.Named,
) string {
	for _, entry := range entries {
		found, _, _ := types.LookupFieldOrMethod(
			types.NewPointer(target),
			true,
			target.Obj().Pkg(),
			entry.Symbol,
		)

		if found != nil {
			return fmt.Sprintf(
				"%s already has %s - the moved method would collide with it or capture its callers",
				qualifiedType(target),
				entry.Symbol,
			)
		}
	}

	return ""
}
