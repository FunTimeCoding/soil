package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/lease"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func (c *Client) Leases(phrase string) ([]*lease.Lease, error) {
	rows, e := searchRows[response.Lease](c, constant.LeaseSearch, phrase)

	if e != nil {
		return nil, e
	}

	return lease.NewSlice(rows), nil
}
