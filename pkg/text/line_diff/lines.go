package line_diff

import (
	"github.com/funtimecoding/soil/pkg/text/constant"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func Lines(
	original string,
	modified string,
) []*Line {
	m := diffmatchpatch.New()
	a, b, lines := m.DiffLinesToChars(original, modified)
	var result []*Line

	for _, d := range m.DiffCharsToLines(m.DiffMain(a, b, false), lines) {
		mark := constant.DiffEqual

		switch d.Type {
		case diffmatchpatch.DiffDelete:
			mark = constant.DiffDelete
		case diffmatchpatch.DiffInsert:
			mark = constant.DiffInsert
		case diffmatchpatch.DiffEqual:
		}

		for _, line := range split(d.Text) {
			result = append(result, New(mark, line))
		}
	}

	return result
}
