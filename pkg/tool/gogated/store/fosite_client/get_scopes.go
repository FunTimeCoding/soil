package fosite_client

import (
	"github.com/ory/fosite"
	"strings"
)

func (c *Client) GetScopes() fosite.Arguments {
	return strings.Split(c.Row.Scopes, " ")
}
