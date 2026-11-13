package search_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/url"
)

func (p *Page) path(
	mode string,
	metadata map[string]string,
) string {
	parameters := url.Values{}
	parameters.Set(webConstant.ParameterQuery, p.query)

	if p.collection != "" {
		parameters.Set(constant.Collection, p.collection)
	}

	if mode != constant.HybridMode {
		parameters.Set(constant.Mode, mode)
	}

	for k, v := range metadata {
		parameters.Set(fmt.Sprintf("%s[%s]", constant.Metadata, k), v)
	}

	return fmt.Sprintf("%s?%s", webConstant.SearchPath, parameters.Encode())
}
