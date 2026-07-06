package connection

import "github.com/funtimecoding/soil/pkg/assistant/message"

// Message type
const (
	Authenticate = "auth"
	Event        = "event"
	Ping         = "ping"
	Pong         = "pong"
	Result       = "result"
	Subscribe    = "subscribe_events"
)

type Subscriber func(*message.Message)
