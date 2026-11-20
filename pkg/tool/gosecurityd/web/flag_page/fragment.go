package flag_page

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/url"
)

func fragment(
	path string,
	name string,
) string {
	return join.Empty(
		path,
		"?",
		url.Values{generativeConstant.ParameterIdentifier: {name}}.Encode(),
	)
}
