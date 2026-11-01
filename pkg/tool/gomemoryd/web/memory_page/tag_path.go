package memory_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"net/url"
)

func TagPath(tag string) string {
	parameters := url.Values{}
	parameters.Set(constant.Tag, tag)

	return fmt.Sprintf("%s?%s", constant.MemoriesPath, parameters.Encode())
}
