package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"golang.org/x/tools/go/packages"
)

func noteUnloaded(
	into unloaded,
	broken []*packages.Package,
	lost string,
) {
	for _, p := range broken {
		if _, known := into[p.ID]; known || len(p.Errors) == 0 {
			continue
		}

		into[p.ID] = result.NewCensusUnloaded(p.ID, p.Errors[0].Error(), lost)
	}
}
