package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"net/url"
)

func pageLink(
	page int,
	tag string,
	memoryType string,
	scope string,
) string {
	params := url.Values{}

	if page > 1 {
		params.Set("page", fmt.Sprintf("%d", page))
	}

	if tag != "" {
		params.Set(constant.Tag, tag)
	}

	if memoryType != "" {
		params.Set(constant.Type, memoryType)
	}

	if scope != "" {
		params.Set(constant.Scope, scope)
	}

	query := params.Encode()

	if query == "" {
		return constant.MemoriesPath
	}

	return fmt.Sprintf("%s?%s", constant.MemoriesPath, query)
}
