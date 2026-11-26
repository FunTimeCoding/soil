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
			return pointer.NewLive(full)
		}

		return deadShort(candidate)
	}

	if relative, inside := pointer.Relative(path, candidate); inside &&
		r.Exists(relative) {
		return pointer.NewLive(relative)
	}

	normalized := pointer.Normalize(candidate)

	if target, found := r.ancestor(path, normalized); found {
		return pointer.NewLive(target)
	}

	first, _, _ := strings.Cut(normalized, stringsConstant.Slash)

	if slices.Contains(constant.ConventionDirectories, first) {
		return pointer.NewConcern(constant.ConventionWordKey)
	}

	if r.Dependency(normalized) || r.Stdlib(normalized) {
		return pointer.NewLive("")
	}

	if strings.HasPrefix(normalized, constant.HomePrefix) {
		return pointer.NewUnchecked(constant.ReasonHome)
	}

	if slices.Contains(constant.ImageRegistries, first) ||
		slices.Contains(r.Registries, first) {
		if r.Literal("", normalized) {
			return pointer.NewLive("")
		}

		return pointer.NewUnchecked(constant.ReasonImage)
	}

	if constant.Address.MatchString(first) {
		return pointer.NewUnchecked(constant.ReasonNetwork)
	}

	if strings.Contains(first, stringsConstant.Dot) {
		return pointer.NewUnchecked(constant.ReasonExternal)
	}

	if full, anchored := r.anchored(r.ImplicitBases, candidate); anchored {
		if r.Exists(full) || r.SiblingExists(full) || r.Ignored(full) {
			return pointer.NewLive(full)
		}

		return deadShort(candidate)
	}

	if len(bases) > 0 {
		return pointer.NewUnchecked(constant.ReasonUnanchored)
	}

	return pointer.NewUnchecked(constant.ReasonUnknown)
}
