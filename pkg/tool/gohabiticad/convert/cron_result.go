package convert

import (
	"github.com/funtimecoding/soil/pkg/habitica/cron"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func CronResult(r *cron.Cron) *server.CronResult {
	result := &server.CronResult{
		RolledOver: r.RolledOver,
		LastCron:   r.LastCron,
	}

	if r.Before != nil {
		result.Before = statsPointer(r.Before)
	}

	if r.After != nil {
		result.After = statsPointer(r.After)
	}

	return result
}
