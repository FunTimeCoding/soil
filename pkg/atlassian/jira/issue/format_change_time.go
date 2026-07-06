package issue

import "github.com/funtimecoding/soil/pkg/time"

func (i *Issue) FormatChangeTime() string {
	return i.ChangeTime().Format(time.DateMinute)
}
