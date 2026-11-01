package history_page

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func isEditable(kind string) bool {
	return kind == constant.Complete ||
		kind == constant.Summarize ||
		kind == constant.Moment
}
