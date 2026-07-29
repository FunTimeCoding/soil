package issue

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func ConvertTime(s string) time.Time {
	return library.Parse(constant.JiraTimeFormat, s)
}
