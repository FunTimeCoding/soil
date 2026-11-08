package face

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/types"
)

func universe() []*fact.Interface {
	return scopeInterfaces(constant.BuiltinPackage, types.Universe)
}
