package line_diff_tester

import "github.com/funtimecoding/soil/pkg/text/line_diff"

func Lines(anchors []*line_diff.Anchor) []int {
	var result []int

	for _, a := range anchors {
		result = append(result, a.Line)
	}

	return result
}
