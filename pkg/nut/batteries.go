package nut

import (
	"context"
	"github.com/funtimecoding/soil/pkg/nut/constant"
	"strings"
)

func (c *Client) Batteries(x context.Context) ([]string, error) {
	lines, e := c.exchange(x, constant.ListBatteries, "")

	if e != nil {
		return nil, e
	}

	var result []string

	for _, l := range lines {
		f := strings.Fields(l)

		if len(f) >= 2 && f[0] == constant.BatteryLine {
			result = append(result, f[1])
		}
	}

	return result, nil
}
