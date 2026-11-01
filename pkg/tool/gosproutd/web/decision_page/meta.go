package decision_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/web/component/card"
)

func meta(
	c *card.Card,
	d *decision.Decision,
) *card.Card {
	c.Identifier(Mark(d.Identifier))

	if frames := frameNames(d); frames != "" {
		c.Meta(frames)
	}

	return c.Meta(join.Empty("pushed ", time.FormatCompact(d.CreatedAt)))
}
