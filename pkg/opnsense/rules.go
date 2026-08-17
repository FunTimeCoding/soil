package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
	"github.com/funtimecoding/soil/pkg/opnsense/rule"
)

func (c *Client) Rules(phrase string) ([]*rule.Rule, error) {
	rows, e := searchRows[response.Rule](c, constant.RuleSearch, phrase)

	if e != nil {
		return nil, e
	}

	return rule.NewSlice(rows), nil
}
