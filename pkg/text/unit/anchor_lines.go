package unit

import "github.com/funtimecoding/soil/pkg/text/line_diff"

func anchorLines(anchors []*line_diff.Anchor) []int {
	var result []int

	for _, a := range anchors {
		result = append(result, a.Line)
	}

	return result
}
