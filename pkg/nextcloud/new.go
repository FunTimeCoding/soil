package nextcloud

import (
	"github.com/funtimecoding/soil/pkg/nextcloud/basic"
	"github.com/funtimecoding/soil/pkg/nextcloud/helper"
	"github.com/funtimecoding/soil/pkg/web/author"
)

func New(
	host string,
	user string,
	password string,
) *Client {
	return &Client{
		basic:  basic.New(host, user, password),
		author: author.New(helper.FileRoot(host, user), user, password),
	}
}
