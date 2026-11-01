package entry_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
)

func Locator(identifier uint) string {
	return fmt.Sprintf("%s/%d", constant.EntryPath, identifier)
}
