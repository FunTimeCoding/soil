package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func resolveSystem(candidate string) *pointer.Resolution {
	if strings.HasPrefix(candidate, constant.RestRoutePrefix) {
		return pointer.NewResolution(constant.VerdictBareSlash)
	}

	result := pointer.NewResolution(constant.VerdictTallied)

	if strings.HasSuffix(candidate, stringsConstant.Slash) {
		result.Reason = constant.ReasonSystem

		return result
	}

	base := candidate[strings.LastIndex(candidate, stringsConstant.Slash)+1:]

	if strings.Contains(base, stringsConstant.Dot) {
		result.Reason = constant.ReasonSystem
	} else {
		result.Reason = constant.ReasonSlash
	}

	return result
}
