package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func clientDetailCard(c *client.Client) gomponents.Node {
	return html.Div(
		html.Class("credential-card"),
		detailRow("Identifier", c.Identifier),
		detailRow("Redirect Locators", c.RedirectLocators),
		detailRow("Grant Types", c.GrantTypes),
		detailRow("Response Types", c.ResponseTypes),
		detailRow("Scopes", c.Scopes),
		detailRow("Auth Method", c.TokenEndpointAuthMethod),
		detailRow("Created", c.CreatedAt.Format("2006-01-02 15:04:05")),
	)
}
