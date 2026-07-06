package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gomonitord",
	"Monitor WebSocket server",
	"gomonitord <address>",
)
