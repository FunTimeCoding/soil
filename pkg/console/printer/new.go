package printer

import "github.com/funtimecoding/soil/pkg/console/status/option"

func New(f *option.Format) *Printer {
	return &Printer{format: f}
}
