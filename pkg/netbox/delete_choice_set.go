package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"net/http"
)

func (c *Client) DeleteChoiceSet(name string) error {
	v, e := c.ChoiceSetByName(name)

	if e != nil {
		return e
	}

	return c.rawRequest(
		http.MethodDelete,
		fmt.Sprintf("%s%d/", constant.ChoiceSetPath, v.Identifier),
		nil,
		nil,
	)
}
