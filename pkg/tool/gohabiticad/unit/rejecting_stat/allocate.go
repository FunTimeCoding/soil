package rejecting_stat

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
)

func (c *Client) Allocate(stat string) (*statistic.Statistic, error) {
	return nil, validation.New("%q is not a valid Stat.", stat)
}
