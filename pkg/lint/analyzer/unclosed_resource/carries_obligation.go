package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
	"strings"
)

func carriesObligation(t types.Type) bool {
	if t == nil {
		return false
	}

	if constant.NonResourceCloseTypes[strings.TrimPrefix(t.String(), "*")] {
		return false
	}

	return isResponse(t) || hasCloseMethod(t)
}
