package nut

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/nut/constant"
	"strings"
)

func (c *Client) Variables(
	x context.Context,
	battery string,
) (map[string]string, error) {
	lines, e := c.exchange(
		x,
		fmt.Sprintf(constant.ListVariables, battery),
		battery,
	)

	if e != nil {
		return nil, e
	}

	result := map[string]string{}

	for _, l := range lines {
		f := strings.SplitN(l, " ", 4)

		if len(f) == 4 && f[0] == constant.VariableLine {
			result[f[2]] = unquote(f[3])
		}
	}

	return result, nil
}
