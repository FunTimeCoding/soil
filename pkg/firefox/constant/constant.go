package constant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"time"
)

const (
	HostEnvironment = "FIREFOX_HOST"
	DefaultHost     = "127.0.0.1:6125"

	NoTitle   = "untitled"
	NoLocator = "about:blank"

	ReadLimit = 64 * 1024 * 1024

	WaitTimeout = 30 * time.Second
	SendTimeout = 60 * time.Second
)

var Format = constant.ColorFormat.Copy()
