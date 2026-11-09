package filter

import "github.com/funtimecoding/soil/pkg/time/window"

func New() *Filter {
	return &Filter{Window: window.Whole()}
}
