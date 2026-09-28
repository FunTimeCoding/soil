package fosite_client

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/client"

func New(row *client.Client) *Client {
	return &Client{Row: row}
}
