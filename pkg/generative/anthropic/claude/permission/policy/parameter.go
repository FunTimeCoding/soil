package policy

import "mvdan.cc/sh/v3/syntax"

func (w *Walk) parameter(p *syntax.ParamExp) (string, bool) {
	if p.Param == nil || p.Excl || p.Length || p.Width || p.Index != nil ||
		p.Slice != nil || p.Repl != nil || p.Exp != nil ||
		p.NestedParam != nil || p.Flags != nil || len(p.Modifiers) > 0 {
		return "", false
	}

	value, okay := w.variables[p.Param.Value]

	return value, okay
}
