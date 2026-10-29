package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"

func New() *Client {
	return &Client{
		StatsValue: &client.Stats{},
		GearValue:  &client.GearResult{Equipped: map[string]string{}},
		ScoreValue: &client.ScoreResult{},
		CronValue:  &client.CronResult{},
	}
}
