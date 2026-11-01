package search_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/url"
)

func Path(
	query string,
	collection string,
	metadata map[string]string,
) string {
	parameters := url.Values{}
	parameters.Set(webConstant.ParameterQuery, query)

	if collection != "" {
		parameters.Set(constant.Collection, collection)
	}

	for k, v := range metadata {
		parameters.Set(fmt.Sprintf("%s[%s]", constant.Metadata, k), v)
	}

	return fmt.Sprintf("%s?%s", webConstant.SearchPath, parameters.Encode())
}
