package conversations

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/block"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func hitBlock(
	b *block.Block,
	terms []string,
	matching bool,
	current bool,
) gomponents.Node {
	class := "message message-user"

	if b.Role == "assistant" {
		class = "message message-assistant"
	}

	if b.Kind != constant.BlockMessage {
		class = join.Space(class, "message-tool")
	}

	if matching {
		class = join.Space(class, "search-hit")
	}

	if current {
		class = join.Space(class, "search-current")
	}

	return html.Div(
		html.ID(join.Empty("block-", b.Identifier)),
		html.Class(class),
		html.Div(
			html.Class("message-role"),
			gomponents.Text(blockLabel(b.Role, b.Kind)),
		),
		html.Div(
			html.Class("message-text"),
			gomponents.Group(highlighted(b.Text, terms)),
		),
	)
}
