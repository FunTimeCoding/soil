package graph

import (
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
	"github.com/funtimecoding/soil/pkg/microsoft/face"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"golang.org/x/oauth2"
)

func New(
	graph *locator.Locator,
	login *locator.Locator,
	tenant string,
	client string,
	s face.TokenStore,
) *Client {
	result := &Client{
		pool: stall_client.New().Build(),
		configuration: &oauth2.Config{
			ClientID: client,
			Scopes:   constant.Scopes,
			Endpoint: oauth2.Endpoint{
				DeviceAuthURL: login.Copy().Path(
					constant.DeviceCodePath,
					tenant,
				).String(),
				TokenURL: login.Copy().Path(
					constant.TokenPath,
					tenant,
				).String(),
				AuthStyle: oauth2.AuthStyleInParams,
			},
		},
		store: s,
	}
	result.requester = requester.New(graph).WithAuthorizer(result)

	return result
}
