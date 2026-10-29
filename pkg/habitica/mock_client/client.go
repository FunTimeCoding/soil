package mock_client

import (
	"github.com/funtimecoding/soil/pkg/habitica/gear"
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
	"github.com/funtimecoding/soil/pkg/habitica/tag"
	"github.com/funtimecoding/soil/pkg/habitica/task"
)

type Client struct {
	Failure error
	Closed  int
	tasks   []*task.Task
	tags    []*tag.Tag
	stats   *statistic.Statistic
	gear    *gear.Gear
}
