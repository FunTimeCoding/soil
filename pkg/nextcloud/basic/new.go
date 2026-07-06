package basic

import "github.com/funtimecoding/soil/pkg/nextcloud/helper"

func New(
	host string,
	user string,
	password string,
) *Client {
	return &Client{
		base:     helper.Base(host),
		fileRoot: helper.FileRoot(host, user),
		user:     user,
		password: password,
	}
}
