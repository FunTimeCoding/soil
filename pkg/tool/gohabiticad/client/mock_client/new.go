package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		StatsValue: &client.Stats{},
		GearValue:  &client.GearResult{Equipped: map[string]string{}},
		ScoreValue: &client.ScoreResult{},
		CronValue:  &client.CronResult{},
	}
	t.Cleanup(result.Close)

	return result
}
