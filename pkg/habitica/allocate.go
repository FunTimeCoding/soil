package habitica

import (
	"github.com/funtimecoding/soil/pkg/habitica/constant"
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
	"net/url"
)

func (c *Client) Allocate(stat string) (*statistic.Statistic, error) {
	var result *statistic.Statistic
	e := c.basic.Post(
		"/user/allocate",
		url.Values{constant.StatParameter: {stat}},
		nil,
		&result,
	)

	return result, e
}
