package history_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func EventPath(
	identifier uint,
	edit bool,
) string {
	result := fmt.Sprintf("%s/%d", constant.HistoryPath, identifier)

	if edit {
		return fmt.Sprintf("%s/%s", result, constant.EditAction)
	}

	return result
}
