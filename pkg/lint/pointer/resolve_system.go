package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func resolveSystem(candidate string) *Resolution {
	if strings.HasPrefix(candidate, constant.RestRoutePrefix) {
		return &Resolution{Verdict: constant.VerdictBareSlash}
	}

	if strings.HasSuffix(candidate, stringsConstant.Slash) {
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonSystem,
		}
	}

	base := candidate[strings.LastIndex(candidate, stringsConstant.Slash)+1:]

	if strings.Contains(base, stringsConstant.Dot) {
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonSystem,
		}
	}

	return &Resolution{
		Verdict: constant.VerdictTallied,
		Reason:  constant.ReasonSlash,
	}
}
