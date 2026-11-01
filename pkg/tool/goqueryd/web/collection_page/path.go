package collection_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"net/url"
	"strconv"
)

func Path(
	name string,
	sourceType string,
	page int,
) string {
	base := fmt.Sprintf("%s/%s", constant.CollectionsPath, url.PathEscape(name))
	parameters := url.Values{}

	if sourceType != "" {
		parameters.Set(constant.SourceType, sourceType)
	}

	if page > 1 {
		parameters.Set(constant.Page, strconv.Itoa(page))
	}

	if len(parameters) == 0 {
		return base
	}

	return fmt.Sprintf("%s?%s", base, parameters.Encode())
}
