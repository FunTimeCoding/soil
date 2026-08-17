package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
	"github.com/funtimecoding/soil/pkg/opnsense/source_nat"
)

func (c *Client) SourceNatRules(phrase string) ([]*source_nat.Rule, error) {
	rows, e := searchRows[response.SourceNatRule](
		c,
		constant.SourceNatSearch,
		phrase,
	)

	if e != nil {
		return nil, e
	}

	return source_nat.NewSlice(rows), nil
}
