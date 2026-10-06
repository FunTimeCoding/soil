package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func (r *Resolver) Resolve(
	path string,
	d *Declared,
	c *Candidate,
) *Resolution {
	span, fragment, found := strings.Cut(c.Span, constant.FragmentSeparator)

	if !found ||
		strings.Contains(c.Span, constant.LocatorSeparator) ||
		(span == "" && !c.Link) {
		return r.resolveSpan(path, d, c)
	}

	result := &Resolution{Verdict: constant.VerdictLive, Target: path}

	if span != "" {
		result = r.resolveSpan(path, d, &Candidate{Span: span, Link: c.Link})
	}

	if result.Verdict != constant.VerdictLive {
		return result
	}

	return r.resolveFragment(result.Target, fragment)
}
