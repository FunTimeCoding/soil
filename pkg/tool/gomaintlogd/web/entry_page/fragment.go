package entry_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
)

func fragment(
	path string,
	identifier uint,
) string {
	return fmt.Sprintf("%s?%s=%d", path, constant.Identifier, identifier)
}
