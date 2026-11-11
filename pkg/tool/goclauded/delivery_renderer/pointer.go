package delivery_renderer

import (
	"fmt"
	"github.com/dustin/go-humanize"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
)

func (r *Renderer) pointer(m *message.Message) string {
	return fmt.Sprintf(
		constant.DeliveryPointer,
		m.Identifier,
		m.FromName,
		m.CreatedAt.In(r.location).Format(timeConstant.HourMinute),
		humanize.Comma(int64(length(m.Body))),
	)
}
