package delivery

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
)

func fragment(
	formatted string,
	m *message.Message,
	room int,
) string {
	total := length(formatted)
	kept := min(
		room-length(
			fmt.Sprintf(
				constant.DeliveryCut,
				humanize.Comma(int64(total)),
				m.Identifier,
			),
		),
		total,
	)

	if kept < constant.DeliveryCutFloor {
		return ""
	}

	return join.Empty(
		string([]rune(formatted)[:kept]),
		fmt.Sprintf(
			constant.DeliveryCut,
			humanize.Comma(int64(total-kept)),
			m.Identifier,
		),
	)
}
