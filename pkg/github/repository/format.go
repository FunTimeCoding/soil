package repository

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/time"
)

func (r *Repository) Format(f *option.Format) string {
	return status.New(f).String(
		r.Owner,
		r.Name,
		r.CreatedAt.Format(time.DateMinute),
	).RawList(r).Format()
}
