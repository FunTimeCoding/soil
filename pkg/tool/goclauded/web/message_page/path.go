package message_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func Path(page int) string {
	if page <= 1 {
		return constant.MessagesPath
	}

	return fmt.Sprintf("%s?%s=%d", constant.MessagesPath, constant.Page, page)
}
