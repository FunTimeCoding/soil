package face

import (
	"github.com/funtimecoding/soil/pkg/habitica/cron"
	"github.com/funtimecoding/soil/pkg/habitica/gear"
	"github.com/funtimecoding/soil/pkg/habitica/request"
	"github.com/funtimecoding/soil/pkg/habitica/score"
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
	"github.com/funtimecoding/soil/pkg/habitica/tag"
	"github.com/funtimecoding/soil/pkg/habitica/task"
)

type HabiticaSource interface {
	Allocate(statistic string) (*statistic.Statistic, error)
	CreateTask(*request.CreateTaskBody) (*task.Task, error)
	Cron() (*cron.Cron, error)
	Equip(key string) error
	Score(
		taskIdentifier string,
		direction string,
	) (*score.Score, error)
	Tags() ([]*tag.Tag, error)
	Tasks(taskType string) ([]*task.Task, error)
	UserGear() (*gear.Gear, error)
	UserStats() (*statistic.Statistic, error)
	MustCreateTask(*request.CreateTaskBody) *task.Task
	MustCron() *cron.Cron
	MustScore(
		taskIdentifier string,
		direction string,
	) *score.Score
	MustTags() []*tag.Tag
	MustTasks(taskType string) []*task.Task
	MustUserStats() *statistic.Statistic
}
