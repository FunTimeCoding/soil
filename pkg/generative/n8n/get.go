package n8n

import (
	"github.com/funtimecoding/soil/pkg/generative/n8n/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Client) Get(path string) string {
	r := web.NewGet(
		locator.New(c.host).Base(constant.Base).Path(path).String(),
	)
	r.Header.Add(constant.TokenHeader, c.token)

	return web.ReadString(web.Send(web.Client(), r))
}
