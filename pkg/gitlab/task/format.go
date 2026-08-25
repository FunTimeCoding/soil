package task

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Task) Format(f *option.Format) string {
	return status.New(f).String(
		string(t.Type),
		t.State,
		t.Body,
	).RawList(t.Raw).Format()
}
