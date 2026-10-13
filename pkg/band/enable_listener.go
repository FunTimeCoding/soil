package band

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
)

func (c *Client) EnableListener() error {
	body, e := c.call(constant.GetAction, constant.RedirectionResource, "", "")

	if e != nil {
		return e
	}

	instance := constant.RedirectionInstance.Find(body)

	if instance == nil {
		return unexpected.Format("band redirection: no instance in response")
	}

	bound := constant.RedirectionOpening.ReplaceAll(
		instance,
		[]byte(
			fmt.Sprintf(
				`<${1}:AMT_RedirectionService xmlns:${1}="%s">`,
				constant.RedirectionResource,
			),
		),
	)
	_, f := c.call(
		constant.PutAction,
		constant.RedirectionResource,
		"",
		string(
			constant.ListenerDisabled.ReplaceAll(bound, []byte("${1}true${2}")),
		),
	)

	return f
}
