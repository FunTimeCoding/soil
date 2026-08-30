package web

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"strings"
)

func sharedNamespace(entries []*board_entry.Entry) string {
	prefix := ""

	for index, entry := range entries {
		slash := strings.Index(entry.Project, constant.Slash)

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
