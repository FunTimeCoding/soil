package constant

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"time"
)

const (
	FirefoxDefaultHost = "127.0.0.1:6125"

	FirefoxSelectorRequired   = "selector is required"
	FirefoxExpressionRequired = "expression is required"
	FirefoxLocatorRequired    = "url is required"
	FirefoxTabsRequired       = "tab_ids is required"
	FirefoxMatchedNothing     = "selector matched nothing: "
	FirefoxNotFoundInTime     = "selector not found within timeout: "
	FirefoxNotConnected       = "extension not connected"

	FirefoxNoTitle   = "untitled"
	FirefoxNoLocator = "about:blank"

	FirefoxReadLimit = 64 * 1024 * 1024

	FirefoxWaitTimeout = 30 * time.Second
	FirefoxSendTimeout = 60 * time.Second
)

var FirefoxFormat = consoleConstant.ColorFormat.Copy()
