package open_webui

import "github.com/funtimecoding/soil/pkg/generative/open_webui/basic"

func New(
	host string,
	token string,
) *Client {
	return &Client{basic: basic.New(host, token)}
}
