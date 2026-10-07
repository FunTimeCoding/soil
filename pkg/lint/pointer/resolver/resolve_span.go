package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func (r *Resolver) resolveSpan(
	path string,
	d *pointer.Declared,
	c *pointer.Candidate,
) *pointer.Resolution {
	if c.Link && !pointer.IsPath(c.Span) {
		return r.resolveLink(path, c.Span)
	}

	switch pointer.Classify(c.Span, r.Roots) {
	case constant.PointerClassLocator:
		return resolveLocator(d.Hosts, c.Span)
	case constant.PointerClassShort:
		return r.resolveShort(path, d.Bases, c.Span)
	case constant.PointerClassCommand:
		return r.resolveCommand(d.Commands, c.Span)
	case constant.PointerClassRoute:
		return r.resolveRoute(d.Bases, c.Span)
	case constant.PointerClassSystem:
		return resolveSystem(c.Span)
	case constant.PointerClassPath:
		result := pointer.NewResolution(constant.VerdictTallied)
		result.Reason = constant.ReasonSystem

		return result
	case constant.PointerClassImport:
		result := pointer.NewResolution(constant.VerdictTallied)
		result.Reason = constant.ReasonImport

		return result
	case constant.PointerClassPattern:
		result := pointer.NewResolution(constant.VerdictTallied)
		result.Reason = constant.ReasonPattern

		return result
	case constant.PointerClassAbsolute:
		return pointer.NewResolution(constant.VerdictAbsolute)
	case constant.PointerClassSibling:
		return r.resolveSibling(path, c.Span)
	case constant.PointerClassSymbol:
		return r.resolveSymbol(c.Span)
	case constant.PointerClassRepository:
		return r.resolveRepository(c.Span)
	}

	return pointer.NewResolution(constant.VerdictLive)
}
