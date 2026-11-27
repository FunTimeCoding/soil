package constant

import "time"

const (
	LadybirdDriverPath = "/Applications/Ladybird.app/Contents/MacOS/WebDriver"

	LadybirdListenFlag   = "--listen-address"
	LadybirdPortFlag     = "--port"
	LadybirdProfileFlag  = "--profile"
	LadybirdHeadlessFlag = "--headless"

	LadybirdStatusRoute  = "/status"
	LadybirdSessionRoute = "/session"

	LadybirdWebSocketCapability = "webSocketUrl"
)

const (
	LadybirdTreeMethod        = "browsingContext.getTree"
	LadybirdNavigateMethod    = "browsingContext.navigate"
	LadybirdCaptureMethod     = "browsingContext.captureScreenshot"
	LadybirdEvaluateMethod    = "script.evaluate"
	LadybirdCallMethod        = "script.callFunction"
	LadybirdActionsMethod     = "input.performActions"
	LadybirdReleaseMethod     = "input.releaseActions"
	LadybirdSubscribeMethod   = "session.subscribe"
	LadybirdCompleteReadiness = "complete"
	LadybirdExceptionResult   = "exception"

	LadybirdNavigationStartedEvent = "browsingContext.navigationStarted"
	LadybirdLoadEvent              = "browsingContext.load"
	LadybirdNavigationFailedEvent  = "browsingContext.navigationFailed"
	LadybirdNavigationAbortedEvent = "browsingContext.navigationAborted"
	LadybirdFragmentNavigatedEvent = "browsingContext.fragmentNavigated"
	LadybirdHistoryUpdatedEvent    = "browsingContext.historyUpdated"

	LadybirdKeySource      = "key"
	LadybirdKeyboardSource = "keyboard"
	LadybirdKeyDown        = "keyDown"
	LadybirdKeyUp          = "keyUp"
	LadybirdStringValue    = "string"
)

const (
	LadybirdSuccessType = "success"
	LadybirdErrorType   = "error"
	LadybirdEventType   = "event"

	LadybirdNoSuchNode      = "no such node"
	LadybirdNoSuchFrame     = "no such frame"
	LadybirdInvalidArgument = "invalid argument"
	LadybirdUnknownCommand  = "unknown command"

	LadybirdNotConnected = "browser not connected"

	LadybirdFixtureSession = "fixture-session"
	LadybirdFixtureAddress = "127.0.0.1:4444"
)

const (
	LadybirdReadLimit = 64 * 1024 * 1024

	LadybirdDriverWait  = 20 * time.Second
	LadybirdDriverPoll  = 100 * time.Millisecond
	LadybirdSendTimeout = 60 * time.Second
)
