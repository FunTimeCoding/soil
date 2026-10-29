package constant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"time"
)

const (
	DefaultHost = "127.0.0.1:6125"

	SelectorRequired   = "selector is required"
	ExpressionRequired = "expression is required"
	LocatorRequired    = "url is required"
	TabsRequired       = "tab_ids is required"
	MatchedNothing     = "selector matched nothing: "
	NotFoundInTime     = "selector not found within timeout: "
	NotConnected       = "extension not connected"

	NoTitle   = "untitled"
	NoLocator = "about:blank"

	ReadLimit = 64 * 1024 * 1024

	WaitTimeout = 30 * time.Second
	SendTimeout = 60 * time.Second
)

var Format = constant.ColorFormat.Copy()
