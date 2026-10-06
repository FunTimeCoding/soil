package reflow

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/yuin/goldmark/v2/ast"
	"strings"
	"unicode"
	"unicode/utf8"
)

func interruption(
	source []byte,
	list ast.Node,
) (int, bool) {
	above, okay := list.PreviousSibling().(ast.BlockNode)

	if !okay || above.Kind() != ast.KindParagraph {
		return 0, false
	}

	lines := above.Source()
	item := firstBlock(list)

	if len(lines) == 0 || item == nil || len(item.Source()) == 0 {
		return 0, false
	}

	last := lines[len(lines)-1]
	start := item.Source()[0].Start
	itemLine := lineStart(source, start)
	between := string(source[last.Start:itemLine])

	if itemLine < last.Stop ||
		len(bytes.TrimSpace(source[last.Stop:itemLine])) != 0 ||
		strings.Count(between, stringsConstant.Unix) > 1 {
		return 0, false
	}

	tail := strings.TrimRight(
		strings.TrimSpace(string(source[last.Start:last.Stop])),
		constant.EmphasisMarkers,
	)

	if strings.HasSuffix(tail, stringsConstant.Colon) {
		return 0, false
	}

	rest := source[start:]

	for len(rest) > 0 {
		r, size := utf8.DecodeRune(rest)

		if !strings.ContainsRune(constant.OpeningMarks, r) {
			return itemLine, unicode.IsLower(r)
		}

		rest = rest[size:]
	}

	return 0, false
}
