package conversations

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/url"
)

func hitsLocator(
	session string,
	query string,
	kinds []string,
	around string,
) string {
	values := url.Values{}
	values.Set(constant.QueryField, query)

	for _, kind := range kinds {
		values.Add(constant.KindField, kind)
	}

	values.Set(constant.Around, around)

	return fmt.Sprintf(
		"/conversations/%s/hits?%s",
		url.PathEscape(session),
		values.Encode(),
	)
}
