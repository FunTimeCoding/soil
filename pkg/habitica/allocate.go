package habitica

import (
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (c *Client) Allocate(stat string) (*statistic.Statistic, error) {
	var result *statistic.Statistic

	return result, c.post(
		join.Empty("/user/allocate?stat=", stat),
		nil,
		&result,
	)
}
