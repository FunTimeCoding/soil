package issue

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	soilTime "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func ConvertTime(s string) time.Time {
	return soilTime.Parse(constant.JiraTimeFormat, s)
}
