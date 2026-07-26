package issue

import "github.com/funtimecoding/soil/pkg/time/constant"

func (i *Issue) FormatChangeTime() string {
	return i.ChangeTime().Format(constant.DateMinute)
}
