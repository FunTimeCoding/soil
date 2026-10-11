package unit

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"

func targetByIdentifier(
	targets []*target.Target,
	identifier string,
) *target.Target {
	for _, t := range targets {
		if t.Identifier == identifier {
			return t
		}
	}

	return nil
}
