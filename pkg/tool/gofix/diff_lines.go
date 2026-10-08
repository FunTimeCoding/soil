package gofix

import "github.com/sergi/go-diff/diffmatchpatch"

func diffLines(
	original []byte,
	modified []byte,
) []*DiffLine {
	m := diffmatchpatch.New()
	a, b, lines := m.DiffLinesToChars(string(original), string(modified))
	var result []*DiffLine

	for _, d := range m.DiffCharsToLines(m.DiffMain(a, b, false), lines) {
		mark := " "

		switch d.Type {
		case diffmatchpatch.DiffDelete:
			mark = "-"
		case diffmatchpatch.DiffInsert:
			mark = "+"
		case diffmatchpatch.DiffEqual:
		}

		for _, line := range splitLines([]byte(d.Text)) {
			result = append(result, &DiffLine{Mark: mark, Text: line})
		}
	}

	return result
}
