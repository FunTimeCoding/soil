package flag_page

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
)

func LookupForm(name string) gomponents.Node {
	return entry.New(constant.LookupPath, "Look up").
		Query().
		Text(generativeConstant.ParameterIdentifier, "CVE", name).
		Required().
		Placeholder("e.g. CVE-2024-3094").
		Inline().
		Node()
}
