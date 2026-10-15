package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"net/http"
)

func (c *Client) DeleteConfigurationContext(name string) error {
	v, e := c.ConfigurationContextByName(name)

	if e != nil {
		return e
	}

	return c.rawRequest(
		http.MethodDelete,
		fmt.Sprintf("%s%d/", constant.ConfigurationContextPath, v.Identifier),
		nil,
		nil,
	)
}
