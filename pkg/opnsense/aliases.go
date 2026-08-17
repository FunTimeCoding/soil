package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/alias"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func (c *Client) Aliases(phrase string) ([]*alias.Alias, error) {
	rows, e := searchRows[response.Alias](c, constant.AliasSearch, phrase)

	if e != nil {
		return nil, e
	}

	return alias.NewSlice(rows), nil
}
