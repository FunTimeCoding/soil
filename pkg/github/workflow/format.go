package workflow

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/time"
)

func (w *Workflow) Format(f *option.Format) string {
	return status.New(f).String(
		w.Name,
		w.State,
		w.CreatedAt.Format(time.DateMinute),
	).RawList(w).Format()
}
