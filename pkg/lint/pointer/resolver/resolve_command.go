package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"slices"
	"strings"
)

func (r *Resolver) resolveCommand(
	declared []string,
	candidate string,
) *pointer.Resolution {
	if slices.Contains(declared, candidate) {
		return pointer.NewResolution(constant.VerdictLive)
	}

	name := strings.TrimPrefix(candidate, stringsConstant.Slash)

	if slices.Contains(constant.HarnessCommands, name) {
		return pointer.NewResolution(constant.VerdictLive)
	}

	for _, p := range pointer.CommandPaths(candidate) {
		if r.Exists(p) || r.SiblingExists(p) {
			return pointer.NewResolution(constant.VerdictLive)
		}
	}

	return pointer.NewResolution(constant.VerdictDead)
}
