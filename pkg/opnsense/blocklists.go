package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/blocklist"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func (c *Client) Blocklists(phrase string) ([]*blocklist.Blocklist, error) {
	rows, e := searchRows[response.Blocklist](
		c,
		constant.BlocklistSearch,
		phrase,
	)

	if e != nil {
		return nil, e
	}

	return blocklist.NewSlice(rows), nil
}
