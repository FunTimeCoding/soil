package report

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/text/line_diff"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
	"strings"
)

func EntryText(e *entry.Entry) []string {
	switch e.Class {
	case constant.ClassBody, constant.ClassSignature, constant.ClassMovedChanged:
		return line_diff.Hunks(
			line_diff.Lines(e.From.Text, e.To.Text),
			constant.TextContext,
		)
	case constant.ClassAdded:
		return strings.Split(e.To.Text, stringsConstant.Unix)
	case constant.ClassRemoved:
		first, _, _ := strings.Cut(e.From.Text, stringsConstant.Unix)

		return []string{first}
	default:
		return nil
	}
}
