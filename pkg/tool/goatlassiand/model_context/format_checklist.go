package model_context

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
)

func formatChecklist(v []*checklist_item.Item) string {
	var lines []string

	for _, i := range v {
		if i.Checked {
			lines = append(lines, join.Empty("+ ", i.Text))
		} else {
			lines = append(lines, join.Empty("- ", i.Text))
		}
	}

	return join.Empty(join.NewLine(lines), separator.Unix)
}
