package history_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/url"
	"strconv"
)

func Path(
	kinds []string,
	page int,
) string {
	parameters := url.Values{}

	for _, k := range kinds {
		parameters.Add(constant.Kind, k)
	}

	if page > 1 {
		parameters.Set(constant.Page, strconv.Itoa(page))
	}

	if len(parameters) == 0 {
		return constant.HistoryPath
	}

	return fmt.Sprintf("%s?%s", constant.HistoryPath, parameters.Encode())
}
