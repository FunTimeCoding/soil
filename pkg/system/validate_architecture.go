package system

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"slices"
)

func ValidateArchitecture(s string) {
	if !slices.Contains(constant.Architectures, s) {
		unexpected.String(s)
	}
}
