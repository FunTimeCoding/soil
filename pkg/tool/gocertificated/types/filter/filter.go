package filter

import "github.com/funtimecoding/soil/pkg/time/window"

type Filter struct {
	Authority string
	Kind      string
	Window    *window.Window
	Revoked   *bool
	Limit     int
}
