package changed

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/git/constant"
)

func (r *Range) String() string {
	switch {
	case r.None:
		return constant.RangeNone
	case r.All:
		return constant.RangeAll
	case r.Staged:
		return constant.RangeStaged
	default:
		return fmt.Sprintf("%s..%s", r.Base, r.Head)
	}
}
