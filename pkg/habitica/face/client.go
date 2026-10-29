package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/habitica/cron"
	"github.com/funtimecoding/soil/pkg/habitica/gear"
	"github.com/funtimecoding/soil/pkg/habitica/request"
	"github.com/funtimecoding/soil/pkg/habitica/score"
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
	"github.com/funtimecoding/soil/pkg/habitica/tag"
	"github.com/funtimecoding/soil/pkg/habitica/task"
)

type Client interface {
	Allocate(stat string) (*statistic.Statistic, error)
	Close()
	CreateTask(b *request.CreateTaskBody) (*task.Task, error)
	Cron() (*cron.Cron, error)
	Equip(key string) error
	Reachable(x context.Context) error
	Score(
		taskIdentifier string,
		direction string,
	) (*score.Score, error)
	Tags() ([]*tag.Tag, error)
	Tasks(taskType string) ([]*task.Task, error)
	UserGear() (*gear.Gear, error)
	UserStats() (*statistic.Statistic, error)
}
