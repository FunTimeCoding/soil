package constant

const (
	HostEnvironment  = "HOME_ASSISTANT_HOST"
	TokenEnvironment = "HOME_ASSISTANT_TOKEN"
)

// Message type
const (
	Authenticate           = "auth"
	AuthenticationRequired = "auth_required"
	AuthenticationSuccess  = "auth_ok"
	AuthenticationInvalid  = "auth_invalid"
	Event                  = "event"
	Ping                   = "ping"
	Pong                   = "pong"
	Result                 = "result"
	Subscribe              = "subscribe_events"
)

const Path = "/api/websocket"
const (
	EventContext = "event"
	TypeKey      = "type"
	RawKey       = "raw"
)
