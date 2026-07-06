package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goquery",
	"Semantic search CLI for goqueryd",
	"goquery [command]",
)

const (
	HostEnvironment = "QUERY_HOST"
	PortEnvironment = "QUERY_PORT"
)
