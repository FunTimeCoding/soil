package census

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/census/reading"

func merge(
	into *reading.Reading,
	from *reading.Reading,
) {
	for k := range from.Calls {
		into.Calls[k] = true
	}

	for k := range from.Fields {
		into.Fields[k] = true
	}

	for k := range from.Keys {
		into.Keys[k] = true
	}

	for k := range from.Packages {
		into.Packages[k] = true
	}

	for l, methods := range from.Uses {
		if into.Uses[l] == nil {
			into.Uses[l] = map[string]bool{}
		}

		for m := range methods {
			into.Uses[l][m] = true
		}
	}

	into.Atomically = into.Atomically || from.Atomically
}
