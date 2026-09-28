package fosite_client

import (
	"github.com/ory/fosite"
	"strings"
)

func (c *Client) GetResponseTypes() fosite.Arguments {
	return strings.Split(c.Row.ResponseTypes, " ")
}
