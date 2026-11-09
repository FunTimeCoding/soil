package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/roster_card"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/component"
	"maragu.dev/gomponents"
)

func ActivityDot(c *roster_card.Card) gomponents.Node {
	return component.ActivityDot(activity(c.Session))
}
