package status

import "github.com/funtimecoding/soil/pkg/console/status/option"

func New(f *option.Format) *Status {
	return &Status{format: f}
}
