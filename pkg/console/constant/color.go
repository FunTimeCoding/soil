package constant

import "github.com/fatih/color"

var (
	Blue    = color.New(color.FgBlue).SprintfFunc()
	Cyan    = color.New(color.FgCyan).SprintfFunc()
	Green   = color.New(color.FgGreen).SprintfFunc()
	Magenta = color.New(color.FgMagenta).SprintfFunc()
	Red     = color.New(color.FgRed).SprintfFunc()
	Yellow  = color.New(color.FgYellow).SprintfFunc()
)
