package constant

import "time"

const (
	LadybirdDriverPath = "/Applications/Ladybird.app/Contents/MacOS/WebDriver"

	ListenFlag   = "--listen-address"
	PortFlag     = "--port"
	ProfileFlag  = "--profile"
	HeadlessFlag = "--headless"

	StatusRoute  = "/status"
	SessionRoute = "/session"

	WebSocketCapability = "webSocketUrl"
)

const (
	TreeMethod       = "browsingContext.getTree"
	NavigateMethod   = "browsingContext.navigate"
	CaptureMethod    = "browsingContext.captureScreenshot"
	EvaluateMethod   = "script.evaluate"
	CallMethod       = "script.callFunction"
	ActionsMethod    = "input.performActions"
	ReleaseMethod    = "input.releaseActions"
	SubscribeMethod  = "session.subscribe"
	CompleteReadiness = "complete"
	ExceptionResult  = "exception"

	NavigationStartedEvent = "browsingContext.navigationStarted"
	LoadEvent              = "browsingContext.load"
	NavigationFailedEvent  = "browsingContext.navigationFailed"
	NavigationAbortedEvent = "browsingContext.navigationAborted"
	FragmentNavigatedEvent = "browsingContext.fragmentNavigated"
	HistoryUpdatedEvent    = "browsingContext.historyUpdated"

	KeySource      = "key"
	KeyboardSource = "keyboard"
	KeyDown        = "keyDown"
	KeyUp          = "keyUp"
	StringValue    = "string"
)

const (
	SuccessType = "success"
	ErrorType   = "error"
	EventType   = "event"

	NoSuchNode      = "no such node"
	NoSuchFrame     = "no such frame"
	InvalidArgument = "invalid argument"
	UnknownCommand  = "unknown command"

	NotConnected = "browser not connected"

	FixtureSession = "fixture-session"
	FixtureAddress = "127.0.0.1:4444"
)

const (
	ReadLimit = 64 * 1024 * 1024

	DriverWait  = 20 * time.Second
	DriverPoll  = 100 * time.Millisecond
	SendTimeout = 60 * time.Second
)
