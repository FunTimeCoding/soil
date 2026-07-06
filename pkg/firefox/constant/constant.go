package constant

import "github.com/funtimecoding/soil/pkg/console/status/option"

const (
	HostEnvironment = "FIREFOX_HOST"
	DefaultHost     = "127.0.0.1:6125"

	NoTitle   = "untitled"
	NoLocator = "about:blank"
)

var Format = option.Color.Copy()
