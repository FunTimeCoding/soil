package repository

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
)

func (r *Repository) Format(f *option.Format) string {
	return status.New(f).String(
		r.Owner,
		r.Name,
		r.CreatedAt.Format(timeConstant.DateMinute),
	).RawList(r).Format()
}
