package constant

import "github.com/funtimecoding/soil/pkg/console/status/option"

var (
	ColorFormat         = option.New().Color()
	ExtendedColorFormat = option.New().Extended().Color()
)
