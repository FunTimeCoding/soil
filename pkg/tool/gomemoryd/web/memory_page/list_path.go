package memory_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"net/url"
	"strconv"
)

func ListPath(
	memoryType string,
	tag string,
	scope string,
	page int,
) string {
	parameters := url.Values{}

	if page > 1 {
		parameters.Set(constant.Page, strconv.Itoa(page))
	}

	if tag != "" {
		parameters.Set(constant.Tag, tag)
	}

	if memoryType != "" {
		parameters.Set(constant.Type, memoryType)
	}

	if scope != "" {
		parameters.Set(constant.Scope, scope)
	}

	if len(parameters) == 0 {
		return constant.MemoriesPath
	}

	return fmt.Sprintf("%s?%s", constant.MemoriesPath, parameters.Encode())
}
