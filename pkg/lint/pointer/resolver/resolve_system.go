package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func resolveSystem(candidate string) *pointer.Resolution {
	if strings.HasPrefix(candidate, constant.RestRoutePrefix) {
		return pointer.NewConcern(constant.BareSlashKey)
	}

	if strings.HasSuffix(candidate, stringsConstant.Slash) {
		return pointer.NewUnchecked(constant.ReasonSystem)
	}

	base := candidate[strings.LastIndex(candidate, stringsConstant.Slash)+1:]

	if strings.Contains(base, stringsConstant.Dot) {
		return pointer.NewUnchecked(constant.ReasonSystem)
	}

	return pointer.NewUnchecked(constant.ReasonSlash)
}
