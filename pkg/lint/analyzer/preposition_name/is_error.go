package preposition_name

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func isError(t types.Type) bool {
	return types.Identical(t, types.Universe.Lookup(constant.ErrorType).Type())
}
