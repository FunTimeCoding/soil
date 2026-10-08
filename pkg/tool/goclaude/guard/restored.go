package guard

import "mvdan.cc/sh/v3/syntax"

func restored(
	d *Edit,
	calls []*syntax.CallExpr,
) bool {
	if d.target == "" {
		return false
	}

	for _, c := range calls {
		if c.Pos().Offset() > d.end && restores(c, d.target) {
			return true
		}
	}

	return false
}
