package pointer

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func (r *Resolver) resolveFragment(
	target string,
	fragment string,
) *Resolution {
	if target == "" {
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonHeading,
		}
	}

	if !strings.HasSuffix(target, library.MarkdownExtension) {
		return &Resolution{
			Verdict: constant.VerdictFragmentTarget,
			Target:  target,
		}
	}

	headings, found := r.Headings(target)

	if !found {
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonHeading,
			Target:  target,
		}
	}

	for _, h := range headings {
		if h.Slug == fragment {
			return &Resolution{Verdict: constant.VerdictLive, Target: target}
		}
	}

	return &Resolution{
		Verdict: constant.VerdictDeadHeading,
		Target:  target,
		Hint:    nearest(fragment, headings),
	}
}
