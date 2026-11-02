package check_entry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
)

func (e *Entry) Format(timestamp bool) string {
	if timestamp {
		return fmt.Sprintf(
			"%s %s %s",
			e.Time.Format(timeConstant.DateMinute),
			e.Level,
			e.Text,
		)
	}

	return key_value.Space(e.Level, e.Text)
}
