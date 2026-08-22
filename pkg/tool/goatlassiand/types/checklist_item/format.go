package checklist_item

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Format(v []*Item) string {
	var lines []string

	for _, i := range v {
		if i.Checked {
			lines = append(lines, join.Empty("+ ", i.Text))
		} else {
			lines = append(lines, join.Empty("- ", i.Text))
		}
	}

	return join.Empty(join.NewLine(lines), constant.Unix)
}
