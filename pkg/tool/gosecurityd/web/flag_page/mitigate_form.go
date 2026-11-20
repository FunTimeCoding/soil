package flag_page

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func MitigateForm(name string) gomponents.Node {
	return entry.New(constant.MitigatePath, "Mitigated").
		Swap(join.Empty("#", constant.FlagsMark), webConstant.SwapOuter).
		Hidden(generativeConstant.ParameterIdentifier, name).
		Text("mitigation", "How", "").
		Placeholder("e.g. upgraded").
		Inline().
		Node()
}
