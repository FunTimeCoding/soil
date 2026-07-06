package chromium

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Client) Tabs() []*tab.Tab {
	var result []*tab.Tab
	notation.MustDecode(
		web.GetString(
			web.InsecureClient(),
			locator.New(
				c.host,
			).Port(c.port).Path(constant.NotationPath).Insecure().String(),
		),
		&result,
		true,
	)

	return result
}
