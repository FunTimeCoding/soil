package stray_comment

import "go/ast"

func isEmptinessMarker(
	c *ast.Comment,
	regions []Region,
) bool {
	if c.Text != "// pass" && c.Text != "// marker" {
		return false
	}

	for _, r := range regions {
		if c.Pos() > r.Start && c.End() < r.End {
			return true
		}
	}

	return false
}
