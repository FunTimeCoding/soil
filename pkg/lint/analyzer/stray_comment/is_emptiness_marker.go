package stray_comment

import "go/ast"

func isEmptinessMarker(
	c *ast.Comment,
	regions []region,
) bool {
	if c.Text != "// pass" && c.Text != "// marker" {
		return false
	}

	for _, r := range regions {
		if c.Pos() > r.From && c.End() < r.To {
			return true
		}
	}

	return false
}
