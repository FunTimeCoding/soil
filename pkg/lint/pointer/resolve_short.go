package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"slices"
	"strings"
)

func (r *Resolver) resolveShort(
	path string,
	bases []string,
	candidate string,
) *Resolution {
	if full, anchored := r.anchored(bases, candidate); anchored {
		if r.Exists(full) || r.SiblingExists(full) || r.Ignored(full) {
			return &Resolution{Verdict: constant.VerdictLive}
		}

		return &Resolution{Verdict: constant.VerdictDead}
	}

	if relative, inside := Relative(path, candidate); inside &&
		r.Exists(relative) {
		return &Resolution{Verdict: constant.VerdictLive}
	}

	normalized := Normalize(candidate)

	if r.ancestorExists(path, normalized) {
		return &Resolution{Verdict: constant.VerdictLive}
	}

	first, _, _ := strings.Cut(normalized, stringsConstant.Slash)

	if slices.Contains(constant.ConventionDirectories, first) {
		return &Resolution{Verdict: constant.VerdictConvention}
	}

	if r.Dependency(normalized) || r.Stdlib(normalized) {
		return &Resolution{Verdict: constant.VerdictLive}
	}

	if strings.HasPrefix(normalized, constant.HomePrefix) {
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonHome,
		}
	}

	if slices.Contains(constant.ImageRegistries, first) ||
		slices.Contains(r.Registries, first) {
		if r.Literal("", normalized) {
			return &Resolution{Verdict: constant.VerdictLive}
		}

		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonImage,
		}
	}

	if constant.Address.MatchString(first) {
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonNetwork,
		}
	}

	if strings.Contains(first, stringsConstant.Dot) {
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonExternal,
		}
	}

	if full, anchored := r.anchored(r.ImplicitBases, candidate); anchored {
		if r.Exists(full) || r.SiblingExists(full) || r.Ignored(full) {
			return &Resolution{Verdict: constant.VerdictLive}
		}

		return &Resolution{Verdict: constant.VerdictDead}
	}

	if len(bases) > 0 {
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonUnanchored,
		}
	}

	return &Resolution{
		Verdict: constant.VerdictTallied,
		Reason:  constant.ReasonUnknown,
	}
}
