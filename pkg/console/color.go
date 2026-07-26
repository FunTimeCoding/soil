package console

import "github.com/fatih/color"

var (
	BlueInstance    = color.New(color.FgBlue)
	CyanInstance    = color.New(color.FgCyan)
	GreenInstance   = color.New(color.FgGreen)
	MagentaInstance = color.New(color.FgMagenta)
	RedInstance     = color.New(color.FgRed)
	YellowInstance  = color.New(color.FgYellow)

	Blue    = BlueInstance.SprintfFunc()
	Cyan    = CyanInstance.SprintfFunc()
	Green   = GreenInstance.SprintfFunc()
	Magenta = MagentaInstance.SprintfFunc()
	Red     = RedInstance.SprintfFunc()
	Yellow  = YellowInstance.SprintfFunc()
)
