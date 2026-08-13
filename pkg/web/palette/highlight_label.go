package palette

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func highlightLabel(
	label string,
	positions []int,
) gomponents.Node {
	if len(positions) == 0 {
		return gomponents.Text(label)
	}

	runes := []rune(label)
	set := make(map[int]bool, len(positions))

	for _, p := range positions {
		set[p] = true
	}

	var nodes []gomponents.Node
	var run []rune
	inMatch := false

	for i, r := range runes {
		matched := set[i]

		if matched != inMatch {
			if len(run) > 0 {
				text := string(run)

				if inMatch {
					nodes = append(nodes, html.Strong(gomponents.Text(text)))
				} else {
					nodes = append(nodes, gomponents.Text(text))
				}

				run = nil
			}

			inMatch = matched
		}

		run = append(run, r)
	}

	if len(run) > 0 {
		text := string(run)

		if inMatch {
			nodes = append(nodes, html.Strong(gomponents.Text(text)))
		} else {
			nodes = append(nodes, gomponents.Text(text))
		}
	}

	return gomponents.Group(nodes)
}
