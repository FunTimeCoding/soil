package memory_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func Path(identifier int64) string {
	return fmt.Sprintf("%s/%d", constant.MemoriesPath, identifier)
}
