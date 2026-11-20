package system_page

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"net/url"
)

func link(name string) string {
	return join.Empty(
		constant.SystemPath,
		"?",
		url.Values{generativeConstant.ParameterName: {name}}.Encode(),
	)
}
