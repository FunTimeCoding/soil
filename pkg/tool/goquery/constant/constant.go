package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goquery",
	"Semantic search CLI for goqueryd",
	"goquery [command]",
)

const (
	HostEnvironment     = "GOQUERY_HOST"
	PortEnvironment     = "GOQUERY_PORT"
	InsecureEnvironment = "GOQUERY_INSECURE"
	TokenEnvironment    = "GOQUERY_TOKEN" // #nosec G101 not a hardcoded secret
)
