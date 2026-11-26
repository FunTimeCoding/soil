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
		return pointer.NewUnchecked(constant.ReasonSystem)
	case constant.PointerClassMedia:
		return resolveMedia(c.Span)
	case constant.PointerClassImport:
		return pointer.NewUnchecked(constant.ReasonImport)
	case constant.PointerClassPattern:
		return pointer.NewUnchecked(constant.ReasonPattern)
	case constant.PointerClassTest:
		return r.resolveTestName(c.Span)
	case constant.PointerClassBareSymbol:
		return pointer.NewUnchecked(constant.ReasonBareSymbol)
	case constant.PointerClassAbsolute:
		return pointer.NewConcern(constant.AbsolutePointerKey)
	case constant.PointerClassSibling:
		return r.resolveSibling(path, c.Span)
	case constant.PointerClassSymbol:
		return r.resolveSymbol(c.Span)
	case constant.PointerClassRepository:
		return r.resolveRepository(c.Span)
	}

	return pointer.NewLive("")
}
