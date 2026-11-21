package flag_page

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/url"
)

func FlagForm(
	values url.Values,
	alert string,
) gomponents.Node {
	return entry.New(constant.FlagPath, "Flag").
		Identifier(constant.FlagFormMark).
		Swap(join.Empty("#", constant.FlagsMark), webConstant.SwapOuter).
		Alert(alert).
		Text(
			generativeConstant.ParameterIdentifier,
			"CVE identifier or remediation item (required)",
			values.Get(generativeConstant.ParameterIdentifier),
		).
		Required().
		Placeholder("e.g. CVE-2024-3094 or upgrade:bookworm:openssl").
		Text(constant.NoteField, "Note", values.Get(constant.NoteField)).
		Placeholder("e.g. runs on the web hosts").
		Node()
}
