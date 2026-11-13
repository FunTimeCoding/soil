package line_diff

import "github.com/funtimecoding/soil/pkg/text/constant"

func Anchors(
	original string,
	modified string,
) []*Anchor {
	var result []*Anchor
	line := 0
	start := 0
	inside := false

	for _, l := range Lines(original, modified) {
		if l.Mark == constant.DiffEqual {
			line++
			inside = false
			result = append(result, NewAnchor(line, false))

			continue
		}

		if !inside {
			inside = true
			start = line + 1
		}

		if l.Mark == constant.DiffDelete {
			line++

			continue
		}

		result = append(result, NewAnchor(max(min(start, line), 1), true))
	}

	return result
}
