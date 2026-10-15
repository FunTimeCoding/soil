package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"net/http"
)

func (c *Client) DeleteCustomField(name string) error {
	v, e := c.CustomFieldByName(name)

	if e != nil {
		return e
	}

	return c.rawRequest(
		http.MethodDelete,
		fmt.Sprintf("%s%d/", constant.CustomFieldPath, v.Identifier),
		nil,
		nil,
	)
}
