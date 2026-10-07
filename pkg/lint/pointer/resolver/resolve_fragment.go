package resolver

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"strings"
)

func (r *Resolver) resolveFragment(
	target string,
	fragment string,
) *pointer.Resolution {
	if target == "" {
		result := pointer.NewResolution(constant.VerdictTallied)
		result.Reason = constant.ReasonHeading

		return result
	}

	if !strings.HasSuffix(target, library.MarkdownExtension) {
		result := pointer.NewResolution(constant.VerdictFragmentTarget)
		result.Target = target

		return result
	}

	headings, found := r.Headings(target)

	if !found {
		result := pointer.NewResolution(constant.VerdictTallied)
		result.Reason = constant.ReasonHeading
		result.Target = target

		return result
	}

	for _, h := range headings {
		if h.Slug == fragment {
			result := pointer.NewResolution(constant.VerdictLive)
			result.Target = target

			return result
		}
	}

	result := pointer.NewResolution(constant.VerdictDeadHeading)
	result.Target = target
	result.Hint = nearest(fragment, headings)

	return result
}
