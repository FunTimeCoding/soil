package imports

import (
	"github.com/funtimecoding/soil/pkg/source/types/import_decision"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
)

func Decide(v *import_view.View) []*import_decision.Decision {
	return append(decideConstants(v), decideCollisions(v)...)
}
