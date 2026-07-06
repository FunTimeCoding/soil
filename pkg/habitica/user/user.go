package user

import (
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
	"github.com/funtimecoding/soil/pkg/habitica/user_items"
)

type User struct {
	NeedsCron bool                 `json:"needsCron"`
	LastCron  string               `json:"lastCron"`
	Stats     *statistic.Statistic `json:"stats"`
	Items     *user_items.Items    `json:"items"`
}
