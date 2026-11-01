package client_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	"github.com/funtimecoding/soil/pkg/web/component/fact"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Detail(c *client.Client) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H1(
				gomponents.Textf(
					"Client %s",
					component.Truncate(c.Identifier, 8),
				),
			),
			fact.New().
				Row("Identifier", c.Identifier).
				Row("Redirect Locators", c.RedirectLocators).
				Row("Grant Types", c.GrantTypes).
				Row("Response Types", c.ResponseTypes).
				Row("Scopes", c.Scopes).
				Row("Auth Method", c.TokenEndpointAuthMethod).
				Row("Created", time.FormatCompact(c.CreatedAt)).
				Node(),
			action.New(
				join.Empty(Path(c.Identifier), constant.DeleteSuffix),
				"Delete Client",
			).
				ButtonClass("secondary").
				Confirm(fmt.Sprintf("Delete client %s?", c.Identifier)).
				Node(),
		},
	)
}
