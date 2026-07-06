package alert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func collect() []*alert.Alert {
	return common.Opsgenie().Open()
}
