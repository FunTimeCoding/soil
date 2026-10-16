package host_mapping

import (
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"slices"
)

func stronger(
	current string,
	next string,
) string {
	if current == "" ||
		slices.Index(constant.HostPrecedence, next) <
			slices.Index(constant.HostPrecedence, current) {
		return next
	}

	return current
}
