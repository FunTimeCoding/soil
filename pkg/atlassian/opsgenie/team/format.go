package team

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Team) Format(f *option.Format) string {
	r := status.New(f).String(t.Name, t.Identifier).RawList(t.Raw)

	if t.Description != "" {
		r.String(t.Description)
	}

	return r.Format()
}
