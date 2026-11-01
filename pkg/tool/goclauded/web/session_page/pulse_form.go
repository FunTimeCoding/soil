package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
)

func (p *Page) pulseForm() gomponents.Node {
	return entry.New(
		ActionPath(p.view.Detail.Identifier, constant.PulseAction),
		"Send",
	).
		Inline().
		Text(constant.Body, "Pulse", "").
		Placeholder("pulse...").
		Checkbox(constant.Immediate, "wake now", false).
		Node()
}
