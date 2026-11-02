package alert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func collect() []*alert.Alert {
	c := common.Opsgenie()
	defer c.Close()

	return c.Open()
}
