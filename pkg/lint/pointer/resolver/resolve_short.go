package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"slices"
	"strings"
)

func (r *Resolver) resolveShort(
	path string,
	bases []string,
	candidate string,
) *pointer.Resolution {
	if full, anchored := r.anchored(bases, candidate); anchored {
		if r.Exists(full) || r.SiblingExists(full) || r.Ignored(full) {
			result := pointer.NewResolution(constant.VerdictLive)
			result.Target = full

			return result
		}

		return pointer.NewResolution(constant.VerdictDead)
	}

	if relative, inside := pointer.Relative(path, candidate); inside &&
		r.Exists(relative) {
		result := pointer.NewResolution(constant.VerdictLive)
		result.Target = relative

		return result
	}

	normalized := pointer.Normalize(candidate)

	if target, found := r.ancestor(path, normalized); found {
		result := pointer.NewResolution(constant.VerdictLive)
		result.Target = target

		return result
	}

	first, _, _ := strings.Cut(normalized, stringsConstant.Slash)

	if slices.Contains(constant.ConventionDirectories, first) {
		return pointer.NewResolution(constant.VerdictConvention)
	}

	if r.Dependency(normalized) || r.Stdlib(normalized) {
		return pointer.NewResolution(constant.VerdictLive)
	}

	if strings.HasPrefix(normalized, constant.HomePrefix) {
		result := pointer.NewResolution(constant.VerdictTallied)
		result.Reason = constant.ReasonHome

		return result
	}

	if slices.Contains(constant.ImageRegistries, first) ||
		slices.Contains(r.Registries, first) {
		if r.Literal("", normalized) {
			return pointer.NewResolution(constant.VerdictLive)
		}

		result := pointer.NewResolution(constant.VerdictTallied)
		result.Reason = constant.ReasonImage

		return result
	}

	if constant.Address.MatchString(first) {
		result := pointer.NewResolution(constant.VerdictTallied)
		result.Reason = constant.ReasonNetwork

		return result
	}

	if strings.Contains(first, stringsConstant.Dot) {
		result := pointer.NewResolution(constant.VerdictTallied)
		result.Reason = constant.ReasonExternal

		return result
	}

	if full, anchored := r.anchored(r.ImplicitBases, candidate); anchored {
		if r.Exists(full) || r.SiblingExists(full) || r.Ignored(full) {
			result := pointer.NewResolution(constant.VerdictLive)
			result.Target = full

			return result
		}

		return pointer.NewResolution(constant.VerdictDead)
	}

	result := pointer.NewResolution(constant.VerdictTallied)

	if len(bases) > 0 {
		result.Reason = constant.ReasonUnanchored
	} else {
		result.Reason = constant.ReasonUnknown
	}

	return result
}
