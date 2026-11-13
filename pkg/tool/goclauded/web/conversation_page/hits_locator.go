package conversation_page

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
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

	return join.Empty(
		panelPath(session),
		stringsConstant.Slash,
		constant.HitsAction,
		stringsConstant.Question,
		values.Encode(),
	)
}
