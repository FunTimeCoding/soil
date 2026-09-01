package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Customer() {
	j := common.Jira()
	console.Format("Information: %+v\n", j.MustCustomerInformation())

	if false {
		for _, i := range j.MustCustomerIssues(true) {
			console.Line(i.Format(constant.ColorFormat))
		}
	}

	if false {
		desks := j.MustDesks()

		for _, e := range desks.Values {
			console.Format("Desk: %+v\n", e)
			types := j.MustRequestTypes(strings.MustToInteger(e.ID), 0)

			for _, t := range types.Values {
				console.Format("  Type: %+v\n", t)
			}
		}
	}

	if false {
		i := j.MustCreateCustomerIssue(
			1,
			4,
			"Test software request",
			"Requesting the software",
		)
		console.Format("Issue: %+v\n", i)
	}
}
