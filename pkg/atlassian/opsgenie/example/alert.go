package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Alert() {
	a := argument.NewSimple("opsgenie-alert")
	a.Boolean(constant.Create, false, "Create alert")
	a.String(constant.User, "", "User email for alert")
	a.String(constant.Text, "", "Alert name")
	a.String(constant.Close, "", "Alert ID")
	a.ParseSimple()
	c := common.Opsgenie()

	if a.GetBoolean(constant.Create) {
		c.Create(a.GetString(constant.User), a.GetString(constant.Text))

		return
	}

	if i := a.GetString(constant.Close); i != "" {
		c.Close(i)

		return
	}

	f := consoleConstant.ExtendedColorFormat.Copy()
	alerts := c.Open()

	for _, a := range alerts {
		console.Line(a.Format(f))
	}

	if len(alerts) == 0 {
		console.Line("No relevant alerts")
	}
}
