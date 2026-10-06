package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
)

func reachReplacing(
	o *option.Fix,
	violations []violation,
	r *output.Results,
) ([]*reachedModule, bool) {
	renames := make(map[string]violation)

	for _, v := range violations {
		if v.fix == "" || !v.object.Exported() {
			continue
		}

		if t, okay := xref.Target(v.object); okay {
			renames[t] = v
		}
	}

	if len(renames) == 0 {
		return nil, true
	}

	var result []*reachedModule

	for _, other := range o.Replacing {
		m, e := reachModule(o.IndexDirectory(), other, renames)

		if e != nil {
			r.AddConcern(
				concern.NewFile(
					"replacing-module",
					fmt.Sprintf(
						"could not load %s, no rename applied: %s",
						other,
						e,
					),
					other,
					false,
				),
			)

			return nil, false
		}

		if m != nil {
			result = append(result, m)
		}
	}

	return result, true
}
