package n8n

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Client) Get(path string) string {
	r := web.NewGet(
		locator.New(c.host).Base(constant.N8nBase).Path(path).String(),
	)
	r.Header.Add(constant.N8nTokenHeader, c.token)

	return web.ReadString(web.Send(web.Client(), r))
}
