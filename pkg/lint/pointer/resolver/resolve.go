package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"strings"
)

func (r *Resolver) Resolve(
	path string,
	d *pointer.Declared,
	c *pointer.Candidate,
) *pointer.Resolution {
	span, fragment, found := strings.Cut(c.Span, constant.FragmentSeparator)

	if !found ||
		strings.Contains(c.Span, constant.LocatorSeparator) ||
		(span == "" && !c.Link) {
		return r.resolveSpan(path, d, c)
	}

	result := pointer.NewResolution(constant.VerdictLive)
	result.Target = path

	if span != "" {
		inner := pointer.NewSpan(span)
		inner.Link = c.Link
		result = r.resolveSpan(path, d, inner)
	}

	if result.Verdict != constant.VerdictLive {
		return result
	}

	return r.resolveFragment(result.Target, fragment)
}
