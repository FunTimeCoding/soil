package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Alert() {
	a := argument.NewSimple("opsgenie-alert")
	a.Boolean(argument.Create, false, "Create alert")
	a.String(argument.User, "", "User email for alert")
	a.String(argument.Text, "", "Alert name")
	a.String(argument.Close, "", "Alert ID")
	a.ParseSimple()
	c := common.Opsgenie()

	if a.GetBoolean(argument.Create) {
		c.Create(
			a.GetString(argument.User),
			a.GetString(argument.Text),
		)

		return
	}

	if i := a.GetString(argument.Close); i != "" {
		c.Close(i)

		return
	}

	f := option.ExtendedColor.Copy()
	alerts := c.Open()

	for _, a := range alerts {
		fmt.Println(a.Format(f))
	}

	if len(alerts) == 0 {
		fmt.Println("No relevant alerts")
	}
}
