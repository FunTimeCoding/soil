package board_page

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"strings"
)

func namespace(entries []*board_entry.Entry) string {
	prefix := ""

	for index, entry := range entries {
		slash := strings.Index(entry.Project, stringsConstant.Slash)

		if slash < 0 {
			return ""
		}

		current := entry.Project[:slash+1]

		if index == 0 {
			prefix = current

			continue
		}

		if current != prefix {
			return ""
		}
	}

	return prefix
}
