package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Alert() {
	a := argument.NewSimple("opsgenie-alert")
	a.Boolean(argumentConstant.Create, false, "Create alert")
	a.String(argumentConstant.User, "", "User email for alert")
	a.String(argumentConstant.Text, "", "Alert name")
	a.String(argumentConstant.Close, "", "Alert ID")
	a.ParseSimple()
	c := common.Opsgenie()

	if a.GetBoolean(argumentConstant.Create) {
		c.Create(
			a.GetString(argumentConstant.User),
			a.GetString(argumentConstant.Text),
		)

		return
	}

	if i := a.GetString(argumentConstant.Close); i != "" {
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
