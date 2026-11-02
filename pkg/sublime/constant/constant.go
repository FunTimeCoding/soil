package constant

import consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"

const (
	HostEnvironment = "SUBLIME_HOST"
	DefaultHost     = "127.0.0.1:6123"
)

var Format = consoleConstant.ColorFormat.Copy()

const (
	NoTitle  = "untitled"
	NoSyntax = "unknown"
)
