package coordination

import "github.com/funtimecoding/soil/pkg/tool/goclauded/connector"

func targetByIdentifier(
	targets []*connector.Target,
	identifier string,
) *connector.Target {
	for _, t := range targets {
		if t.Identifier == identifier {
			return t
		}
	}

	return nil
}
