package message

import (
	"fmt"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
)

func (m *Message) FormatChat() string {
	return fmt.Sprintf(
		"%s %s: %s",
		m.Create.Format(timeConstant.DateMinute),
		m.From,
		m.Text,
	)
}
