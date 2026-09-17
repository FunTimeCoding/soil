package pointer

import "github.com/funtimecoding/soil/pkg/lint/constant"

func (r *Resolver) Resolve(
	path string,
	bases []string,
	c *Candidate,
) *Resolution {
	if c.Link && !IsPath(c.Span) {
		return r.resolveLink(path, c.Span)
	}

	switch Classify(c.Span, r.Roots) {
	case constant.PointerClassShort:
		return r.resolveShort(path, bases, c.Span)
	case constant.PointerClassCommand:
		return r.resolveCommand(c.Span)
	case constant.PointerClassRoute:
		return r.resolveRoute(bases, c.Span)
	case constant.PointerClassSystem:
		return resolveSystem(c.Span)
	case constant.PointerClassPath:
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonSystem,
		}
	case constant.PointerClassImport:
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonImport,
		}
	case constant.PointerClassPattern:
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonPattern,
		}
	case constant.PointerClassAbsolute:
		return &Resolution{Verdict: constant.VerdictAbsolute}
	case constant.PointerClassSibling:
		return r.resolveSibling(path, c.Span)
	case constant.PointerClassSymbol:
		return r.resolveSymbol(c.Span)
	case constant.PointerClassRepository:
		return r.resolveRepository(c.Span)
	}

	return &Resolution{Verdict: constant.VerdictLive}
}
