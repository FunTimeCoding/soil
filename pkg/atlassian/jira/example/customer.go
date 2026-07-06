package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Customer() {
	j := common.Jira()
	fmt.Printf("Information: %+v\n", j.MustCustomerInformation())

	if false {
		for _, i := range j.MustCustomerIssues(true) {
			fmt.Println(i.Format(option.Color))
		}
	}

	if false {
		desks := j.MustDesks()

		for _, e := range desks.Values {
			fmt.Printf("Desk: %+v\n", e)
			types := j.MustRequestTypes(strings.MustToInteger(e.ID), 0)

			for _, t := range types.Values {
				fmt.Printf("  Type: %+v\n", t)
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
		fmt.Printf("Issue: %+v\n", i)
	}
}
