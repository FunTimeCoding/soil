package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"

type Client struct {
	TaskList    []*client.Task
	TagList     []*client.Tag
	StatsValue  *client.Stats
	GearValue   *client.GearResult
	ScoreValue  *client.ScoreResult
	CronValue   *client.CronResult
	CreatedList []*client.Task
	Failure     error
	Closed      int
}
