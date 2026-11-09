package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func testSpan(
	r *resolver.Resolver,
	moves *move_set.Set,
	span string,
) (string, bool) {
	name := strings.TrimSuffix(span, constant.CallSuffix)

	for _, m := range moves.Symbols {
		if m.Name != name {
			continue
		}

		if r.TestElsewhere(name, m.Directory) {
			return "", false
		}

		return join.Empty(m.NewName, span[len(name):]), true
	}

	return "", true
}
