package chromium

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Tabs() ([]*tab.Tab, error) {
	var result []*tab.Tab

	if e := c.requester.Notation(request.Get(constant.NotationPath), &result); e != nil {
		return nil, e
	}

	return result, nil
}
