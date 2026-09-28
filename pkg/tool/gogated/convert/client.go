package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"github.com/funtimecoding/soil/pkg/tool/gogated/types/summary"
	"strings"
)

func Client(c *client.Client) *summary.Summary {
	return summary.New(
		c.Identifier,
		strings.Split(c.RedirectLocators, "\n"),
		strings.Fields(c.Scopes),
		strings.Fields(c.GrantTypes),
		strings.Fields(c.ResponseTypes),
		c.TokenEndpointAuthMethod,
		c.Public,
		c.CreatedAt,
	)
}
