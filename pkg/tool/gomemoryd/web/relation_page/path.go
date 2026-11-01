package relation_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"net/url"
)

func Path(relationType string) string {
	if relationType == "" {
		return constant.RelationsPath
	}

	parameters := url.Values{}
	parameters.Set(constant.Type, relationType)

	return fmt.Sprintf("%s?%s", constant.RelationsPath, parameters.Encode())
}
