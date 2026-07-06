package habitica

import "github.com/funtimecoding/soil/pkg/habitica/cron"

func (c *Client) Cron() (*cron.Cron, error) {
	before, e := c.user()

	if e != nil {
		return cron.Stub(), e
	}

	if !before.NeedsCron {
		result := cron.New()
		result.LastCron = before.LastCron

		return result, nil
	}

	if f := c.postDiscard("/cron"); f != nil {
		return cron.Stub(), f
	}

	after, g := c.user()

	if g != nil {
		return cron.Stub(), g
	}

	result := cron.New()
	result.RolledOver = true
	result.LastCron = after.LastCron
	result.Before = before.Stats
	result.After = after.Stats

	return result, nil
}
