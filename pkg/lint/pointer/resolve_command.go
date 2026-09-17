package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"slices"
	"strings"
)

func (r *Resolver) resolveCommand(candidate string) *Resolution {
	name := strings.TrimPrefix(candidate, stringsConstant.Slash)

	if slices.Contains(constant.HarnessCommands, name) {
		return &Resolution{Verdict: constant.VerdictLive}
	}

	for _, p := range CommandPaths(candidate) {
		if r.Exists(p) || r.SiblingExists(p) {
			return &Resolution{Verdict: constant.VerdictLive}
		}
	}

	return &Resolution{Verdict: constant.VerdictDead}
}
