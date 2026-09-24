package unit

import (
	"github.com/funtimecoding/soil/pkg/measure"
	"github.com/funtimecoding/soil/pkg/measure/option"
	"github.com/funtimecoding/soil/pkg/measure/registry"
	"github.com/funtimecoding/soil/pkg/measure/result"
)

func scan(root string) *result.Result {
	o := option.New()
	o.Paths = []string{root}

	return measure.Scan(o, registry.NewDefault())
}
