package resolver

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"strings"
)

func (r *Resolver) resolveFragment(
	target string,
	fragment string,
) *pointer.Resolution {
	if target == "" {
		return pointer.NewUnchecked(constant.ReasonHeading)
	}

	if !strings.HasSuffix(target, soil.MarkdownExtension) {
		result := pointer.NewConcern(constant.FragmentTargetKey)
		result.Target = target

		return result
	}

	headings, found := r.Headings(target)

	if !found {
		result := pointer.NewUnchecked(constant.ReasonHeading)
		result.Target = target

		return result
	}

	for _, h := range headings {
		if h.Slug == fragment {
			return pointer.NewLive(target)
		}
	}

	result := pointer.NewConcern(constant.DeadHeadingKey)
	result.Target = target
	result.Hint = nearest(fragment, headings)

	return result
}
