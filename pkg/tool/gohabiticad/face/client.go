package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
)

type Client interface {
	face.Reachable
	Address() string
	AllocateStat(
		x context.Context,
		stat string,
	) (*client.Stats, error)
	CreateTask(
		x context.Context,
		taskType string,
		text string,
		notes string,
	) (*client.Task, error)
	EquipGear(
		x context.Context,
		key string,
	) (*client.GearResult, error)
	GetGear(x context.Context) (*client.GearResult, error)
	GetStats(x context.Context) (*client.Stats, error)
	GetTags(x context.Context) ([]*client.Tag, error)
	GetTasks(
		x context.Context,
		taskType string,
	) ([]*client.Task, error)
	RunCron(x context.Context) (*client.CronResult, error)
	ScoreTask(
		x context.Context,
		identifier string,
		direction string,
	) (*client.ScoreResult, error)
	Close()
}
