package basic

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

// Reference: https://developer.atlassian.com/cloud/confluence/rest/v2
func New(
	root *locator.Locator,
	user string,
	token string,
) *Client {
	base := root.Copy().Base(constant.ConfluenceBase)

	return &Client{
		requester: newRequester(base, user, token),
		old: newRequester(
			root.Copy().Base(constant.ConfluenceOldBase),
			user,
			token,
		),
		root: root,
		base: base,
	}
}
