package cron

import "github.com/funtimecoding/soil/pkg/habitica/statistic"

type Cron struct {
	RolledOver bool                 `json:"rolled_over"`
	LastCron   string               `json:"last_cron"`
	Before     *statistic.Statistic `json:"before,omitempty"`
	After      *statistic.Statistic `json:"after,omitempty"`
}
