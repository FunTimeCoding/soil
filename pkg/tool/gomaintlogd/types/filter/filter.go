package filter

import "github.com/funtimecoding/soil/pkg/time/window"

type Filter struct {
	System  string
	Service string
	User    string
	Window  *window.Window
	Limit   int
}
