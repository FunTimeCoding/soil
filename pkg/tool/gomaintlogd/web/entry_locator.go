package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
)

func entryLocator(identifier uint) string {
	return fmt.Sprintf("%s/%d", constant.EntryPath, identifier)
}
