package message

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/time/constant"
)

func (m *Message) FormatChat() string {
	return fmt.Sprintf(
		"%s %s: %s",
		m.Create.Format(constant.DateMinute),
		m.From,
		m.Text,
	)
}
