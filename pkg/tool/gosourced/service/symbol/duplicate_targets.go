package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol_rename"
)

func duplicateTargets(renames []*symbol_rename.Rename) string {
	sources := map[string]bool{}
	targets := map[string]string{}

	for _, n := range renames {
		source := memberName(n.Receiver, n.OldName)
		target := memberName(n.Receiver, n.NewName)

		if sources[source] {
			return fmt.Sprintf("%s is renamed twice", source)
		}

		if previous, found := targets[target]; found {
			return fmt.Sprintf(
				"%s and %s both rename to %s",
				previous,
				source,
				target,
			)
		}

		sources[source] = true
		targets[target] = source
	}

	return ""
}
