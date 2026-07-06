package alertmanager

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
)

func (c *Client) SilenceExists(name string) (bool, error) {
	_, e := c.SilenceByRule(name)

	if e != nil {
		if errors.Is(e, constant.ErrorNotFound) {
			return false, nil
		}

		return false, e
	}

	return true, nil
}
