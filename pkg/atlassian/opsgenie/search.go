package opsgenie

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	rawAlert "github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) Search(
	query string,
	a ...any,
) []*alert.Alert {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	var result []*alert.Alert
	var start int
	p := c.AlertOption()

	for {
		page, e := c.userClient.Alert.List(
			c.context,
			&rawAlert.ListAlertRequest{
				Limit:  constant.PageLimit,
				Offset: start,
				Query:  query,
			},
		)
		errors.PanicOnError(e)
		result = append(
			result,
			alert.NewSlice(page.Alerts, p, c.Verbose)...,
		)

		if len(page.Alerts) < constant.PageLimit {
			break
		}

		start += constant.PageLimit
	}

	return c.processor().Process(result)
}
